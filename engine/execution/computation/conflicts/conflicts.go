package conflicts

import (
	"fmt"
	"github.com/onflow/flow-go/fvm/state"
	"github.com/onflow/flow-go/model/flow"
	"sync"
)

type Conflicts struct {
	transactionViews map[flow.Identifier]*state.View
	dependencyGraph  Graph
	conflicts        []flow.Identifier
	txChannel        chan Message
	lock             sync.RWMutex
}

type Message struct {
	TransactionID   flow.Identifier
	TransactionView *state.View
	TxIndex         uint32
}

func (c *Conflicts) String() string {
	return c.dependencyGraph.String()
}

func NewConflicts(txCount int) *Conflicts {
	return &Conflicts{
		transactionViews: nil,
		dependencyGraph:  newGraph(),
		conflicts:        []flow.Identifier{},
		txChannel:        make(chan Message, txCount),
	}
}

func (c *Conflicts) AddTransaction(msg Message) {
	c.txChannel <- msg
}

func (c *Conflicts) TransactionCount() int {
	return len(c.transactionViews)
}

func (c *Conflicts) Run() {
	for msg := range c.txChannel {
		c.addTransaction(msg.TransactionID, msg.TransactionView)
	}
}

func (c *Conflicts) Close() {
	close(c.txChannel)
}

func (c *Conflicts) addTransaction(txID flow.Identifier, txView *state.View) error {
	c.lock.Lock()
	if c.transactionViews == nil {
		c.transactionViews = make(map[flow.Identifier]*state.View)
	}
	c.transactionViews[txID] = txView
	c.dependencyGraph.addNode(&Node{transaction: txID})
	c.lock.Unlock()
	return nil
}

// Graph adapted from https://flaviocopes.com/golang-data-structure-graph/
type Graph struct {
	nodes []*Node
	edges map[Node][]*Node
	lock  sync.RWMutex
}

func newGraph() Graph {
	return Graph{
		nodes: []*Node{},
		edges: nil,
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

func (g *Graph) addNode(n *Node) {
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}
func (g *Graph) addDirectedEdge(n1, n2 *Node) {
	g.lock.Lock()
	if g.edges == nil {
		g.edges = make(map[Node][]*Node)
	}
	g.edges[*n1] = append(g.edges[*n1], n2)
	g.lock.Unlock()
}

type Node struct {
	transaction flow.Identifier
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.transaction)
}
