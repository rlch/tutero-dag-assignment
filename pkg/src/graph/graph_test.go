package graph

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNodes(t *testing.T) {
	graph := Graph{}
	tbls := []struct {
		input map[Node][]Node
		want  []Node
	}{
		{map[Node][]Node{}, []Node{}},
		{map[Node][]Node{"A": {}, "B": {}}, []Node{"A", "B"}},
		{map[Node][]Node{"A": {"C", "D"}}, []Node{"A"}}, // C, D are not nodes; they are invalid.
	}
	for _, tbl := range tbls {
		graph.adjacenyList = tbl.input
		got := graph.Nodes()
		sort.Slice(got, func(i, j int) bool {
			return got[i] < got[j]
		})
		if !cmp.Equal(got, tbl.want) {
			t.Errorf("Nodes() = %v, want %v", got, tbl.want)
		}
	}
}

func TestAddNodeOk(t *testing.T) {
	graph := Graph{}
	node := Node("A")
	if err := graph.AddNode(node); err != nil {
		t.Error(err)
	}
	if _, ok := graph.adjacenyList[node]; !ok {
		t.Errorf("AddNode() did not add node to adjacenyList")
	}
}

func TestAddNodeAlreadyExists(t *testing.T) {
	graph := Graph{}
	node := Node("A")
	if err := graph.AddNode(node); err != nil {
		t.Error(err)
	}
	if err := graph.AddNode(node); err == nil {
		t.Error("expected an error")
	}
}

func TestAddEdgeOk(t *testing.T) {
	graph := Graph{}
	u := Node("A")
	v := Node("B")
	graph.AddNode(u)
	graph.AddNode(v)
	if err := graph.AddEdge(u, v); err != nil {
		t.Error(err)
	}
	if edges := graph.adjacenyList[u]; len(edges) != 1 {
		t.Errorf("AddEdge() did not add edge to adjacenyList")
	}

	another := Node("C")
	graph.AddNode(another)
	if edges := graph.adjacenyList[u]; len(edges) != 1 {
		t.Errorf("AddEdge() did not add another edge to adjacenyList")
	}
}

func TestAddEdgeMultiple(t *testing.T) {
	graph := Graph{}
	u := Node("A")
	for i := 0; i < 10; i++ {
		graph.AddEdge(u, Node(fmt.Sprint(i)))
	}
	if edges := graph.adjacenyList[u]; len(edges) != 10 {
		t.Errorf("AddEdge() resulted in incorrect number of edges; got %d, want %d", len(edges), 10)
	}
}

func TestAddEdgeCreatesNodes(t *testing.T) {
	graph := Graph{}
	u := Node("A")
	v := Node("B")
	if err := graph.AddEdge(u, v); err != nil {
		t.Error(err)
	}
	if edges := graph.adjacenyList[u]; len(edges) != 1 {
		t.Errorf("AddEdge() did not add edge to adjacenyList")
	}
}

func TestAddEdgeCycleDirect(t *testing.T) {
	graph := Graph{}
	u := Node("A")
	v := Node("B")
	graph.AddEdge(u, v)
	if err := graph.AddEdge(v, u); err == nil {
		t.Error("expected err to be non-nil")
	}
}

func TestAddEdgeCycleIndirect(t *testing.T) {
	graph := Graph{}
	u := Node("A")
	v := Node("B")
	w := Node("C")
	graph.AddEdge(u, w)
	if err := graph.AddEdge(v, u); err != nil {
		t.Error(err)
	}
	if err := graph.AddEdge(w, v); err == nil {
		t.Error("expected err to be non-nil")
	}
}

func TestAddEdgeAlreadyExists(t *testing.T) {
	graph := Graph{}
	u := Node("A")
	v := Node("B")
	graph.AddEdge(u, v)
	if err := graph.AddEdge(u, v); err == nil {
		t.Error("expected err to be non-nil")
	}
}

func TestRemoveNode(t *testing.T) {
	graph := Graph{}
	tbls := []struct {
		remove Node
		input  map[Node][]Node
		want   map[Node][]Node
	}{
		{"B", map[Node][]Node{"A": {"B"}, "B": {}}, map[Node][]Node{"A": {}}},
		{"B", map[Node][]Node{"A": {"B"}, "B": {"C", "D"}, "D": {"F"}}, map[Node][]Node{"A": {}, "D": {"F"}}},
	}
	for _, tbl := range tbls {
		graph.adjacenyList = tbl.input
		err := graph.RemoveNode(tbl.remove)
		if err != nil {
			t.Error(err)
		}
		if !cmp.Equal(graph.adjacenyList, tbl.want) {
			t.Errorf("Parents() = %v, want %v", graph.adjacenyList, tbl.want)
		}
	}

}

func TestAdjacencyList(t *testing.T) {
	graph := Graph{}
	canonical := map[Node][]Node{
		"A": {"B", "C"},
		"B": {"C"},
		"C": {},
	}
	graph.adjacenyList = canonical
	got := graph.AdjacencyList()
	if !cmp.Equal(got, canonical) {
		t.Errorf("AdjacencyList() = %v, want %v", got, canonical)
	}
	got["C"] = append(got["C"], "D")
	if len(graph.adjacenyList["C"]) != 0 {
		t.Errorf("expected AdjacencyList() to return a copy of the adjacency list")
	}
}

func TestChildren(t *testing.T) {
	graph := Graph{}
	tbls := []struct {
		of    Node
		input map[Node][]Node
		want  []Node
	}{
		{"A", map[Node][]Node{"A": {"B"}}, []Node{"B"}},
		{"A", map[Node][]Node{"A": {"B"}, "B": {"C", "D"}}, []Node{"B", "C", "D"}},
	}
	for _, tbl := range tbls {
		graph.adjacenyList = tbl.input
		got := graph.Children(tbl.of)
		sort.Slice(got, func(i, j int) bool {
			return got[i] < got[j]
		})
		if !cmp.Equal(got, tbl.want) {
			t.Errorf("Children() = %v, want %v", got, tbl.want)
		}
	}
}

func TestParents(t *testing.T) {
	graph := Graph{}
	tbls := []struct {
		of    Node
		input map[Node][]Node
		want  []Node
	}{
		{"B", map[Node][]Node{"A": {"B"}}, []Node{"A"}},
		{"F", map[Node][]Node{"A": {"B"}, "B": {"C", "D"}, "D": {"F"}}, []Node{"A", "B", "D"}},
	}
	for _, tbl := range tbls {
		graph.adjacenyList = tbl.input
		got := graph.Parents(tbl.of)
		sort.Slice(got, func(i, j int) bool {
			return got[i] < got[j]
		})
		if !cmp.Equal(got, tbl.want) {
			t.Errorf("Parents() = %v, want %v", got, tbl.want)
		}
	}
}

// TODO: add more test-cases
func TestBFS(t *testing.T) {
	graph := Graph{}
	tbls := []struct {
		source Node
		input  map[Node][]Node
		want   []Node
	}{
		{
			"A",
			map[Node][]Node{
				"A": {"B", "C"},
				"C": {"D"},
				"B": {"F", "G"},
			},
			[]Node{"A", "B", "C", "F", "G", "D"},
		},
		{
			"C",
			map[Node][]Node{
				"A": {"B", "C"},
				"C": {"D"},
				"B": {"F", "G"},
			},
			[]Node{"C", "D"},
		},
	}
	for _, tbl := range tbls {
		graph.adjacenyList = tbl.input
		order := []Node{}
		if err := graph.BreadthFirstSearch(tbl.source, func(node Node) error {
			order = append(order, node)
			return nil
		}); err != nil {
			t.Error(err)
		}
		if !cmp.Equal(order, tbl.want) {
			t.Errorf("BreadthFirstSearch() = %v, want %v", order, tbl.want)
		}
	}
}

func BenchmarkRandom(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Random()
		if err != nil {
			b.Error(err)
		}
	}
}
