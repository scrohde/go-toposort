package toposort

import "testing"

func index(s []string, v string) int {
	for i, s := range s {
		if s == v {
			return i
		}
	}
	return -1
}

type Edge struct {
	From string
	To   string
}

func TestDuplicatedNode(t *testing.T) {
	graph := NewGraph(2)
	graph.AddNode("a")
	if graph.AddNode("a") {
		t.Errorf("not raising duplicated node error")
	}

}

func TestRemoveNotExistEdge(t *testing.T) {
	graph := NewGraph(0)
	if graph.RemoveEdge("a", "b") {
		t.Errorf("not raising not exist edge error")
	}
}

func TestWikipedia(t *testing.T) {
	graph := NewGraph(8)
	graph.AddNodes("2", "3", "5", "7", "8", "9", "10", "11")

	edges := []Edge{
		{"7", "8"},
		{"7", "11"},

		{"5", "11"},

		{"3", "8"},
		{"3", "10"},

		{"11", "2"},
		{"11", "9"},
		{"11", "10"},

		{"8", "9"},
	}

	for _, e := range edges {
		graph.AddEdge(e.From, e.To)
	}

	result, ok := graph.ToposortStable()
	if !ok {
		t.Errorf("closed path detected in no closed pathed graph")
	}

	for _, e := range edges {
		if i, j := index(result, e.From), index(result, e.To); i > j {
			t.Errorf("dependency failed: not satisfy %v(%v) > %v(%v)", e.From, i, e.To, j)
		}
	}
}

func TestCycle(t *testing.T) {
	graph := NewGraph(3)
	graph.AddNodes("1", "2", "3")

	graph.AddEdge("1", "2")
	graph.AddEdge("2", "3")
	graph.AddEdge("3", "1")

	_, ok := graph.Toposort()
	if ok {
		t.Errorf("closed path not detected in closed pathed graph")
	}
}
