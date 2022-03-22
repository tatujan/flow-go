package conflicts

import (
	"fmt"
	"github.com/onflow/flow-go/model/flow"
	"sync"
)

type Conflicts struct {
	transactions            []Transaction
	totalRegisterTouchSet   map[flow.RegisterID][]Transaction
	dependencyGraph         Graph
	conflictingTransactions map[flow.Identifier]Transaction
	txChannel               chan Transaction
	lock                    sync.RWMutex
}

type Transaction struct {
	TransactionID    flow.Identifier
	CollectionID     flow.Identifier
	RegisterTouchSet []flow.RegisterID
	TxIndex          uint32
}

func (t *Transaction) String() string {
	return t.TransactionID.String()
}

func (c *Conflicts) String() string {
	return c.dependencyGraph.String()
}

func NewConflicts(txCount int) *Conflicts {
	return &Conflicts{
		transactions:            make([]Transaction, txCount),
		totalRegisterTouchSet:   make(map[flow.RegisterID][]Transaction),
		dependencyGraph:         newGraph(),
		conflictingTransactions: make(map[flow.Identifier]Transaction),
		txChannel:               make(chan Transaction, txCount),
	}
}

func (c *Conflicts) StoreTransaction(tx Transaction) {
	c.txChannel <- tx
}

func (c *Conflicts) AddConflictingTransaction(tx Transaction, conflicts []Transaction) {
	txID := tx.TransactionID
	c.conflictingTransactions[txID] = tx
	g := c.dependencyGraph
	if !g.Contains(txID) {
		g.AddNode(txID)
	}
	for _, conflict := range conflicts {
		conflictID := conflict.TransactionID
		if !g.Contains(conflictID) {
			g.AddNode(conflictID)
		}
		g.AddDirectedEdge(conflictID, txID)
	}
}

func (c *Conflicts) InTransactionConflicts(tx Transaction) bool {
	_, inMap := c.conflictingTransactions[tx.TransactionID]
	return inMap
}

func (c *Conflicts) TransactionCount() int {
	return len(c.transactions)
}

func (c *Conflicts) Run() {
	for tx := range c.txChannel {
		c.transactions[tx.TxIndex] = tx
		c.transactions[tx.TxIndex] = tx
	}
	c.buildDependencyGraph()
}

func (c *Conflicts) buildDependencyGraph() {
	for _, tx := range c.transactions {
		conflicts := c.collectConflicts(tx)
		if len(conflicts) > 0 {
			c.AddConflictingTransaction(tx, conflicts)
		}
	}
}

func (c *Conflicts) collectConflicts(tx Transaction) []Transaction {
	var conflicts []Transaction
	for _, register := range tx.RegisterTouchSet {
		// for each register tx touches, check if it has already been touched by previous transactionNode
		registerTouches, registerExists := c.totalRegisterTouchSet[register]
		if registerExists {
			// if it has been touched previously, check if the transactionNode that previously touched it are from a different collection or are already in the list of conflicts
			for _, transaction := range registerTouches {
				if transaction.CollectionID != tx.CollectionID || c.InTransactionConflicts(transaction) {
					conflicts = append(conflicts, transaction)
				}
			}
		}
	}
	return conflicts
}

func (c *Conflicts) Close() {
	close(c.txChannel)
}

// Graph adapted from https://flaviocopes.com/golang-data-structure-graph/
type Graph struct {
	transactionNode map[flow.Identifier]*Node
	nodes           []*Node
	edges           map[Node][]*Node
	lock            sync.RWMutex
}

func newGraph() Graph {
	return Graph{
		transactionNode: make(map[flow.Identifier]*Node),
		nodes:           []*Node{},
		edges:           nil,
	}
}

func (g *Graph) String() string {
	g.lock.RLock()
	str := ""
	if len(g.nodes) > 0 {
		for i := 0; i < len(g.nodes); i++ {
			str += g.nodes[i].String()
			edges := g.edges[*g.nodes[i]]
			if len(edges) > 0 {
				str += " -> "
				for j := 0; j < len(edges); j++ {
					str += edges[j].String() + " "
				}
			}
			str += "\n"
		}
	} else {
		str += "Empty graph."
	}
	return str
}

func (g *Graph) AddNode(txID flow.Identifier) {
	g.lock.Lock()
	node := &Node{transaction: txID}
	g.transactionNode[txID] = node
	g.nodes = append(g.nodes, node)
	g.lock.Unlock()
}
func (g *Graph) AddDirectedEdge(tx1, tx2 flow.Identifier) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	n1 := g.transactionNode[tx1]
	n2 := g.transactionNode[tx2]
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.lock.Unlock()
}

func (g *Graph) Contains(txID flow.Identifier) bool {
	_, inMap := g.transactionNode[txID]
	return inMap
}

type Node struct {
	transaction flow.Identifier
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.transaction)
}
