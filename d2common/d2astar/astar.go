package d2astar

import (
	"container/heap"
	"fmt"
	"sync"
)

var nodePool *sync.Pool
var nodeMapPool *sync.Pool
var priorityQueuePool *sync.Pool
func init()  {
	nodePool = &sync.Pool {
		New: func()interface{} {
			return &node{}
		},
	}

	nodeMapPool = &sync.Pool {
		New: func()interface{} {
			return make(nodeMap, 128)
		},
	}

	priorityQueuePool = &sync.Pool {
		New: func()interface{} {
			return priorityQueue{}
		},
	}
}

// astar is an A* pathfinding implementation.

// Pather is an interface which allows A* searching on arbitrary objects which
// can represent a weighted graph.
type Pather interface {
	// PathNeighbors returns the direct neighboring nodes of this node which
	// can be pathed to.
	PathNeighbors() []Pather
	// PathNeighborCost calculates the exact movement cost to neighbor nodes.
	PathNeighborCost(to Pather) float64
	// PathEstimatedCost is a heuristic method for estimating movement costs
	// between non-adjacent nodes.
	PathEstimatedCost(to Pather) float64
}

// node is a wrapper to store A* data for a Pather node.
type node struct {
	pather Pather
	cost   float64
	rank   float64
	parent *node
	open   bool
	closed bool
	index  int
}

func (n *node) reset ()  {
	n.pather = nil
	n.cost = 0
	n.rank = 0
	n.parent = nil
	n.open = false
	n.closed = false
	n.index = 0
}

// nodeMap is a collection of nodes keyed by Pather nodes for quick reference.
type nodeMap map[Pather]*node

// get gets the Pather object wrapped in a node, instantiating if required.
func (nm nodeMap) get(p Pather) *node {
	n, ok := nm[p]
	if !ok {
		n = nodePool.Get().(*node)
		n.pather = p
		nm[p] = n
	}
	return n
}

// Path calculates a short path and the distance between the two Pather nodes.
//
// If no path is found, found will be false.
func Path(from, to Pather, maxCost float64) (path []Pather, distance float64, found bool) {
	// Quick escape for inaccessible areas.
	toNeighbors := to.PathNeighbors()
	if len(toNeighbors) == 0 {
		return nil, 0, false
	}

	nm := nodeMapPool.Get().(nodeMap)
	nq := priorityQueuePool.Get().(priorityQueue)
	defer func() {
		for k, v := range nm {
			v.reset()
			nodePool.Put(v)
			delete(nm, k)
		}

		nq = nq[0:0]
		nodeMapPool.Put(nm)
		priorityQueuePool.Put(nq)
	}()

	heap.Init(&nq)
	fromNode := nm.get(from)
	fromNode.open = true
	heap.Push(&nq, fromNode)
	for {
		if nq.Len() == 0 {
			// There's no path, return found false.
			return
		}
		current := heap.Pop(&nq).(*node)
		current.open = false
		current.closed = true

		if current == nm.get(to) {
			// Found a path to the goal.
			p := make([]Pather, 0, 16)
			curr := current
			for curr != nil {
				p = append(p, curr.pather)
				curr = curr.parent
			}
			return p, current.cost, true
		}

		for _, neighbor := range current.pather.PathNeighbors() {
			cost := current.cost + current.pather.PathNeighborCost(neighbor)
			if cost > maxCost {
				fmt.Println("Canceling path")
				continue
			}

			neighborNode := nm.get(neighbor)
			if cost < neighborNode.cost {
				if neighborNode.open {
					heap.Remove(&nq, neighborNode.index)
				}
				neighborNode.open = false
				neighborNode.closed = false
			}
			if !neighborNode.open && !neighborNode.closed {
				neighborNode.cost = cost
				neighborNode.open = true
				neighborNode.rank = cost + neighbor.PathEstimatedCost(to)
				neighborNode.parent = current
				heap.Push(&nq, neighborNode)
			}
		}
	}
}
