package conflicts

import (
	"errors"
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
		conflicts := c.getConflicts(tx)
		if !conflicts.Empty() {
			c.addToConflictGraph(tx, conflicts)
		}
	}
}

// getConflicts returns a list of previous transactions that conflict with transaction tx. This
func (c *Conflicts) getConflicts(tx Transaction) TransactionSet {
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

// addToConflictGraph takes a conflicts.Transaction type and a list of other Transactions it conflicts with and
// adds those conflicts to the dependency graph
func (c *Conflicts) addToConflictGraph(tx Transaction, conflicts TransactionSet) {
	txID := tx.TransactionID
	// store conflict transaction value using the txID as key
	c.conflictingTransactions[txID] = tx
	g := c.dependencyGraph
	// if the tx is not already in graph, add node
	if !g.Contains(txID) {
		g.addNode(txID)
	}
	// for each transaction that conflicts with tx, add a directed edge from that conflict to tx
	for _, tx := range conflicts.IterableMap() {
		conflictID := tx.TransactionID
		// if the conflict does not already exist in the graph, add it.
		if !g.Contains(conflictID) {
			g.addNode(conflictID)
		}
		g.addDirectedEdge(conflictID, txID)
	}
}

// Graph is a hashmap composed of identifiers as the nodes (keys) and a list of identifiers as edges (values)
type Graph struct {
	graph map[flow.Identifier][]flow.Identifier
	lock  sync.RWMutex
}

func newGraph() Graph {
	return Graph{
		graph: make(map[flow.Identifier][]flow.Identifier),
	}
}

func (g *Graph) Contains(txID flow.Identifier) bool {
	_, inMap := g.graph[txID]
	return inMap
}

func (g *Graph) String() string {
	g.lock.RLock()
	str := ""
	if len(g.graph) > 0 {
		for node, edges := range g.graph {
			str += node.String()
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

func (g *Graph) addNode(txID flow.Identifier) error {
	g.lock.Lock()
	var err error
	// check if node already exists
	if !g.Contains(txID) {
		// if not, add it
		g.graph[txID] = make([]flow.Identifier, 0)
	} else {
		err = errors.New("Node already exists.")
	}
	g.lock.Unlock()
	return err

}

func (g *Graph) deleteNode(txID flow.Identifier) error {
	g.lock.Lock()
	var err error
	if g.Contains(txID) {
		//delete the node
		delete(g.graph, txID)
		// delete node from all edge lists
		for node, edges := range g.graph {
			g.graph[node] = deleteEdgeIfExists(txID, edges)
		}
	} else {
		err = errors.New("Node does not exist.")
	}
	g.lock.Unlock()
	return err

}

// addDirectedEdge adds an edge from tx1 to tx2 by appending tx2 to the list of edges associated with tx1
func (g *Graph) addDirectedEdge(tx1, tx2 flow.Identifier) error {
	g.lock.Lock()
	var err error
	// check if nodes already exist. If not return error
	if !g.Contains(tx1) || !g.Contains(tx2) {
		err = errors.New("Cannot add edge to a node that does not exist.")
	} else {
		edges := g.graph[tx1]
		// check if destination tx2 already exists in list of edges (index >= 0)
		if indexOf(tx2, edges) < 0 {
			// if not, add it
			g.graph[tx1] = append(edges, tx2)
		} else {
			err = errors.New("Edge already exists.")
		}
	}
	g.lock.Unlock()
	return err
}

func (g *Graph) removeDirectedEdge(tx1, tx2 flow.Identifier) error {
	g.lock.Lock()
	var err error
	// check if nodes already exist. If not return error
	if !g.Contains(tx1) || !g.Contains(tx2) {
		err = errors.New("Cannot remove edge from a node that does not exist.")
	} else {
		edges := g.graph[tx1]
		// check if edge already exists, index >= 0
		if index := indexOf(tx2, edges); index >= 0 {
			// if exists, remove edge
			g.graph[tx1] = unorderedDelete(index, edges)
		} else {
			// otherwise, return error
			err = errors.New("Edge does not exist.")
		}
	}
	g.lock.Unlock()
	return err
}

// indexOf checks a list of tx for membership. If exists, index of tx is returned, else -1.
func indexOf(tx flow.Identifier, edges []flow.Identifier) int {
	for index, edge := range edges {
		if tx.String() == edge.String() {
			return index
		}
	}
	return -1
}

// deleteEdgeIfExists takes a tx and a list of edges and removes that tx from the list of edges if it is in the list.
// returns the list of edges, with tx removed if it was found.
func deleteEdgeIfExists(destinationNode flow.Identifier, edges []flow.Identifier) []flow.Identifier {
	if index := indexOf(destinationNode, edges); index >= 0 {
		return unorderedDelete(index, edges)
	} else {
		return edges
	}
}

// unorderedDelete is a constant time delete for an element of a slice identified by index. Slice is reordered.
// return a new slice with index element removed.
func unorderedDelete(index int, edges []flow.Identifier) []flow.Identifier {
	last := len(edges) - 1
	edges[index] = edges[last]
	return edges[:last]
}

// Following code adapted from https://www.davidkaya.com/sets-in-golang/

// TransactionSet stores a set of transactions that are unique via their ID string representation.
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
