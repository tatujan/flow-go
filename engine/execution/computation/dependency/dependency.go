package dependency

import (
	"fmt"
	"github.com/onflow/flow-go/engine/execution/state/delta"
	"sync"
)

package dependency

import (
"fmt"
"github.com/onflow/flow-go/engine/execution/state/delta"
"github.com/onflow/flow-go/model/flow"
"sync"
)

type Dependency struct {
	transactionViews map[flow.Identifier]*delta.View
	dependencyGraph  Graph
	conflicts        []flow.Identifier
	lock             sync.RWMutex
}

func (d *Dependency) String() string {
	return d.dependencyGraph.String()
}

func NewDependency() *Dependency {
	return &Dependency{
		transactionViews: nil,
		dependencyGraph:  newGraph(),
		conflicts:        []flow.Identifier{},
	}
}

func (d *Dependency) AddTransaction(txID flow.Identifier, txView *delta.View) error {
	d.lock.Lock()
	if d.transactionViews == nil {
		d.transactionViews = make(map[flow.Identifier]*delta.View)
	}
	d.transactionViews[txID] = txView
	d.dependencyGraph.addNode(&Node{transaction: txID})
	d.lock.Unlock()
	return nil
}

func (d *Dependency) TransactionCount() int {
	return len(d.transactionViews)
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

