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

// StoreTransaction is the public function that enables adding transactions to the conflict aggregator
func (c *Conflicts) StoreTransaction(tx Transaction) {
	c.txChannel <- tx
}

// TransactionCount returns the number of transactions the conflict object is aware of
func (c *Conflicts) TransactionCount() int {
	return len(c.transactions)
}

// TransactionCount returns the number of transactions the conflict object is aware of
func (c *Conflicts) ConflictCount() int {
	return len(c.conflictingTransactions)
}

// InTransactionConflicts returns true if the tx is in c.conflictingTransactions
func (c *Conflicts) InTransactionConflicts(tx Transaction) bool {
	_, inMap := c.conflictingTransactions[tx.TransactionID]
	return inMap
}

// Run receives transactions and adds them to the list of transactions. When finished receiving transactions, it builds
// the dependency graph.
func (c *Conflicts) Run() {
	for tx := range c.txChannel {
		c.transactions[tx.TxIndex] = tx
	}
	c.buildDependencyGraph()
}

func (c *Conflicts) Close() {
	close(c.txChannel)
}

// buildDependencyGraph iterates over all added transactions one at a time, aggregates conflicts with previous
// transactions,
func (c *Conflicts) buildDependencyGraph() {
	for _, tx := range c.transactions {
		conflicts := c.collectConflicts(tx)
		if !conflicts.Empty() {
			c.addConflictingTransaction(tx, conflicts)
		}
	}
}

// collectConflicts returns a list of previous transactions that conflict with transaction tx. This
func (c *Conflicts) collectConflicts(tx Transaction) TransactionSet {
	var conflicts = NewTransactionSet()
	for _, register := range tx.RegisterTouchSet {
		registerTouches, registerExists := c.totalRegisterTouchSet[register]
		if registerExists {
			// if it has been touched previously, check if the transactionNode that previously touched it are from a
			// different collection or are already in the list of conflicts
			for _, transaction := range registerTouches {
				if transaction.CollectionID != tx.CollectionID || c.InTransactionConflicts(transaction) {
					// if so, that transaction is a conflict. Add to list of conflicts with tx
					conflicts.Add(transaction)
				}
			}
			// add current tx to the list of Transactions that have touched register
			c.totalRegisterTouchSet[register] = append(c.totalRegisterTouchSet[register], tx)
		} else {
			// if register has not been touched yet, initialize the list of transactions with current tx
			c.totalRegisterTouchSet[register] = []Transaction{tx}
		}
	}
	return conflicts
}

// addConflictingTransaction takes a conflicts.Transaction type and a list of other Transactions it conflicts with and
// adds those conflicts to the dependency graph
func (c *Conflicts) addConflictingTransaction(tx Transaction, conflicts TransactionSet) {
	txID := tx.TransactionID
	// store conflict transaction value using the txID as key
	c.conflictingTransactions[txID] = tx
	g := c.dependencyGraph
	// if the tx is not already in graph, add node
	if !g.Contains(txID) {
		g.AddNode(txID)
	}
	// for each transaction that conflicts with tx, add a directed edge from that conflict to tx
	for _, tx := range conflicts.IterableMap() {
		conflictID := tx.TransactionID
		// if the conflict does not already exist in the graph, add it.
		if !g.Contains(conflictID) {
			g.AddNode(conflictID)
		}
		g.AddDirectedEdge(conflictID, txID)
	}
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

// Following code adapted from https://www.davidkaya.com/sets-in-golang/

// TransactionSet stores transactions via their ID string representation
type TransactionSet struct {
	m map[string]Transaction
}

func NewTransactionSet() TransactionSet {
	s := TransactionSet{}
	s.m = make(map[string]Transaction)
	return s
}

func (s *TransactionSet) Add(tx Transaction) {
	value := tx.TransactionID.String()
	s.m[value] = tx
}

func (s *TransactionSet) Remove(tx Transaction) {
	id := tx.TransactionID.String()
	delete(s.m, id)
}

func (s *TransactionSet) Contains(tx Transaction) bool {
	id := tx.TransactionID.String()
	_, c := s.m[id]
	return c
}

func (s *TransactionSet) Empty() bool {
	return len(s.m) == 0
}

func (s *TransactionSet) IterableMap() map[string]Transaction {
	return s.m
}
