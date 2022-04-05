package conflicts

import (
	"errors"
	"github.com/onflow/flow-go/model/flow"
	"sync"
)

type Conflicts struct {
	transactions          []Transaction
	totalRegisterTouchSet map[flow.RegisterID][]Transaction
	dependencyGraph       map[flow.Identifier][]flow.Identifier
	txChannel             chan Transaction
	lock                  sync.RWMutex
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

func NewConflicts(txCount int) *Conflicts {
	return &Conflicts{
		transactions:          make([]Transaction, txCount),
		totalRegisterTouchSet: make(map[flow.RegisterID][]Transaction),
		dependencyGraph:       make(map[flow.Identifier][]flow.Identifier),
		txChannel:             make(chan Transaction, txCount),
	}
}

func (c *Conflicts) String() string {
	c.lock.RLock()
	str := ""
	if len(c.dependencyGraph) > 0 {
		for node, edges := range c.dependencyGraph {
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
		str += "Empty dependency graph."
	}
	return str
}

// StoreTransaction is the public function that enables adding transactions to the conflict aggregator
func (c *Conflicts) StoreTransaction(tx Transaction) {
	c.txChannel <- tx
}

// TransactionCount returns the number of transactions the conflict object is aware of
func (c *Conflicts) TransactionCount() int {
	return len(c.transactions)
}

// ConflictCount returns the number of transactions the conflict object is aware of
// conflicts are defined as nodes that have a directed edge to another node.
func (c *Conflicts) ConflictCount() int {
	count := 0
	for _, edges := range c.dependencyGraph {
		if len(edges) > 0 {
			count++
		}
	}
	return count
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
	// iterate over transactions in order
	for i := 0; i < len(c.transactions); i++ {
		tx := c.transactions[i]
		conflicts := c.determineConflicts(tx)
		if !conflicts.Empty() {
			c.addToConflictGraph(tx.TransactionID, conflicts)
		}
	}
}

// determineConflicts returns a list of previous transactions that conflict with transaction tx. This
func (c *Conflicts) determineConflicts(tx Transaction) TransactionSet {
	var conflicts = NewTransactionSet()
	for _, register := range tx.RegisterTouchSet {
		registerTouches, registerExists := c.totalRegisterTouchSet[register]
		if registerExists {
			// if it has been touched previously, check if the transaction that previously touched it are from a
			// different collection or are already in the list of conflicts
			for _, transaction := range registerTouches {
				if transaction.CollectionID != tx.CollectionID || c.isConflict(transaction.TransactionID) {
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
func (c *Conflicts) addToConflictGraph(txID flow.Identifier, conflicts TransactionSet) {
	// if the tx is not already in graph, add node
	if !c.containsNode(txID) {
		c.addNode(txID)
	}
	// for each transaction that conflicts with tx, add a directed edge from that conflict to tx
	for _, conflict := range conflicts.IterableMap() {
		// if the conflict does not already exist in the graph, add it.
		cID := conflict.TransactionID
		if !c.containsNode(cID) {
			c.addNode(cID)
		}
		c.addDirectedEdge(cID, txID)
	}
}

// a transaction is a conflict if it exists as a node in the dependency graph and has edges leading to other nodes.
func (c *Conflicts) isConflict(txID flow.Identifier) bool {
	edges, inMap := c.dependencyGraph[txID]
	return inMap && len(edges) > 0
}

// ************************************************************
// Graph Operations
// ************************************************************

func (c *Conflicts) containsNode(txID flow.Identifier) bool {
	_, inMap := c.dependencyGraph[txID]
	return inMap
}

func (c *Conflicts) addNode(txID flow.Identifier) error {
	c.lock.Lock()
	var err error
	// check if node already exists
	if !c.containsNode(txID) {
		// if not, add it
		c.dependencyGraph[txID] = make([]flow.Identifier, 0)
	} else {
		err = errors.New("Node already exists.")
	}
	c.lock.Unlock()
	return err
}

func (c *Conflicts) deleteNode(txID flow.Identifier) error {
	c.lock.Lock()
	var err error
	if c.containsNode(txID) {
		//delete the node
		delete(c.dependencyGraph, txID)
		// delete node from all edge lists
		for node, edges := range c.dependencyGraph {
			c.dependencyGraph[node] = deleteEdgeIfExists(txID, edges)
		}
	} else {
		err = errors.New("Node does not exist.")
	}
	c.lock.Unlock()
	return err

}

// addDirectedEdge adds an edge from tx1 to tx2 by appending tx2 to the list of edges associated with tx1
func (c *Conflicts) addDirectedEdge(tx1, tx2 flow.Identifier) error {
	c.lock.Lock()
	var err error
	// check if nodes already exist. If not return error
	if !c.containsNode(tx1) || !c.containsNode(tx2) {
		err = errors.New("Cannot add edge to a node that does not exist.")
	} else {
		edges := c.dependencyGraph[tx1]
		// check if destination tx2 already exists in list of edges (index >= 0)
		if indexOfNodeInEdges(tx2, edges) < 0 {
			// if not, add it
			c.dependencyGraph[tx1] = append(edges, tx2)
		} else {
			err = errors.New("Edge already exists.")
		}
	}
	c.lock.Unlock()
	return err
}

func (c *Conflicts) removeDirectedEdge(tx1, tx2 flow.Identifier) error {
	c.lock.Lock()
	var err error
	// check if nodes already exist. If not return error
	if !c.containsNode(tx1) || !c.containsNode(tx2) {
		err = errors.New("Cannot remove edge from a node that does not exist.")
	} else {
		edges := c.dependencyGraph[tx1]
		// check if edge already exists, index >= 0
		if index := indexOfNodeInEdges(tx2, edges); index >= 0 {
			// if exists, remove edge
			c.dependencyGraph[tx1] = unorderedDelete(index, edges)
		} else {
			// otherwise, return error
			err = errors.New("Edge does not exist.")
		}
	}
	c.lock.Unlock()
	return err
}

// indexOfNodeInEdges checks a list of tx for membership. If exists, index of tx is returned, else -1.
func indexOfNodeInEdges(tx flow.Identifier, edges []flow.Identifier) int {
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
	if index := indexOfNodeInEdges(destinationNode, edges); index >= 0 {
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

// ************************************************************
// Transaction Set Type and Operations
// ************************************************************

// Following code adapted from https://www.davidkaya.com/sets-in-golang/

// TransactionSet stores a set of transactions that are unique via their ID string representation.
type TransactionSet struct {
	m map[flow.Identifier]Transaction
}

func NewTransactionSet() TransactionSet {
	s := TransactionSet{}
	s.m = make(map[flow.Identifier]Transaction)
	return s
}

func (s *TransactionSet) Add(tx Transaction) {
	s.m[tx.TransactionID] = tx
}

func (s *TransactionSet) Remove(tx Transaction) {
	delete(s.m, tx.TransactionID)
}

func (s *TransactionSet) Contains(tx Transaction) bool {
	_, c := s.m[tx.TransactionID]
	return c
}

func (s *TransactionSet) Empty() bool {
	return len(s.m) == 0
}

func (s *TransactionSet) IterableMap() map[flow.Identifier]Transaction {
	return s.m
}
