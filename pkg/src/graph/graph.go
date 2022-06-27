package graph

import (
	"fmt"
	"math/rand"
	"time"
)

type Node string

type Graph struct {
	adjacenyList map[Node][]Node
}

type RandomOptions struct {
	// How fat the DAG is
	MinPerRank int
	// How fat the DAG is
	MaxPerRank int
	// How tall the DAG is
	MinRanks int
	// How tall the DAG is
	MaxRanks int
	// Chance of having an edge
	Percent float32
}

// Random generates a random DAG.
// Algorithm translated from: https://stackoverflow.com/questions/12790337/generating-a-random-dag
func Random(options ...func(*RandomOptions)) (*Graph, error) {
	graph := Graph{}
	rand.Seed(time.Now().UnixNano())
	opts := RandomOptions{
		MinPerRank: 1,
		MaxPerRank: 5,
		MinRanks:   3,
		MaxRanks:   5,
		Percent:    0.3,
	}
	for _, opt := range options {
		opt(&opts)
	}
	randLerp := func(min, max int) int {
		return min + int(rand.Float32()*float32(max-min))
	}

	ranks := randLerp(opts.MinRanks, opts.MaxRanks)
	nodes := 0
	for i := 0; i < ranks; i++ {
		newNodes := randLerp(opts.MinPerRank, opts.MaxPerRank)
		for j := 0; j < nodes; j++ {
			for k := 0; k < newNodes; k++ {
				if rand.Float32() < opts.Percent {
					u := Node(fmt.Sprint(j))
					v := Node(fmt.Sprint(k + nodes))
					if err := graph.AddEdge(u, v); err != nil {
						return nil, err
					}
				}
			}
		}
		nodes += newNodes
	}

	return &graph, nil
}

// Nodes returns the nodes in the graph in no particular order.
func (g Graph) Nodes() []Node {
	nodes := make([]Node, len(g.adjacenyList))
	i := 0
	for node := range g.adjacenyList {
		nodes[i] = node
		i++
	}
	return nodes
}

// AdjacencyList returns an unmodifiable copy of the adjacency list of the graph.
func (g Graph) AdjacencyList() map[Node][]Node {
	cp := make(map[Node][]Node)
	for k, v := range g.adjacenyList {
		cp[k] = make([]Node, len(v))
		for i, node := range v {
			cp[k][i] = node
		}
	}
	return cp
}

// AddNode adds node to the graph with no edges, returning an error if the node already exists.
func (g *Graph) AddNode(node Node) error {
	if g.adjacenyList == nil {
		g.adjacenyList = make(map[Node][]Node)
	}
	if _, ok := g.adjacenyList[node]; ok {
		return fmt.Errorf("attempted to add node %s to graph, but it already exists", node)
	}
	g.adjacenyList[node] = []Node{}
	return nil
}

// AddEdge adds a directed from u -> v
// If u or v do not exist in the graph, they are added to the graph.
// If u -> v already exists, an error is returned.
// If the creation of u -> v results in a cycle (there exists some v -> ... -> u), an error is returned.
func (g *Graph) AddEdge(u, v Node) error {
	if g.adjacenyList == nil {
		g.adjacenyList = make(map[Node][]Node)
	}
	edges, ok := g.adjacenyList[u]
	if !ok {
		g.adjacenyList[u] = []Node{}
	}
	if _, ok := g.adjacenyList[v]; !ok {
		// v does not exist; v -> u is impossible.
		g.adjacenyList[v] = []Node{}
	} else {
		// u and v exist; v -> u is possible.
		if cycle := g.BreadthFirstSearch(v, func(node Node) error {
			if node == u {
				return fmt.Errorf("a cycle was detected when adding %s -> %s", u, v)
			}
			return nil
		}); cycle != nil {
			return cycle
		}
	}
	for _, node := range edges {
		if node == v {
			return fmt.Errorf("%s -> %s already exists", u, v)
		}
	}
	g.adjacenyList[u] = append(edges, v)
	return nil
}

// RemoveNodes removes a node u from the graph, as well as all u -> v and v -> u edges, for all nodes v.
// Returns an error if u was not in the graph.
func (g *Graph) RemoveNode(u Node) error {
	if _, ok := g.adjacenyList[u]; !ok {
		return fmt.Errorf("attempted to remove node %s from graph, but it does not exist", u)
	}
	delete(g.adjacenyList, u)
	for i, edges := range g.adjacenyList {
		for j, edge := range edges {
			if edge == u {
				g.adjacenyList[i] = append(g.adjacenyList[i][:j], g.adjacenyList[i][j+1:]...)
				break
			}
		}
	}
	return nil
}

// DepthFirstSearch performs a depth-first search of the graph, starting at the given node, executing
// `do` for each node visited. If `do` returns an error, the BFS is stopped and the error is returned.
func (g Graph) BreadthFirstSearch(start Node, do func(node Node) error) error {
	visited := map[Node]bool{start: true}
	queue := []Node{start}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		if sentinel := do(cur); sentinel != nil {
			return sentinel
		}
		for _, next := range g.adjacenyList[cur] {
			if !visited[next] {
				queue = append(queue, next)
				visited[next] = true
			}
		}
	}
	return nil
}

// Children walks the graph, returning all recursive children of the given node.
func (g Graph) Children(of Node) []Node {
	children := []Node{}
	g.BreadthFirstSearch(of, func(node Node) error {
		if node != of {
			children = append(children, node)
		}
		return nil
	})
	return children
}

// Parents walks the graph, returning all recursive children of the given node.
func (g Graph) Parents(of Node) []Node {
	visited := map[Node]bool{}
	queue := []Node{of}
	cur := of
	for len(queue) > 0 {
		cur = queue[0]
		queue = queue[1:]
		for next, edges := range g.adjacenyList {
			if visited[next] {
				continue
			}
			for _, to := range edges {
				if to == cur {
					visited[next] = true
					queue = append(queue, next)
					break
				}
			}
		}
	}
	parents := []Node{}
	for node := range visited {
		parents = append(parents, node)
	}
	return parents
}

// TopologicalSort returns a linear ordering of the graph's nodes; guaranteeing that for
// every edge u -> v, u comes before v in the ordering.
func (g Graph) TopologicalSort() ([]Node, error) {
	//* You may wish to implement this!
	panic("not implemented")
}

//* You may implement more methods for Graph!
