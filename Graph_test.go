package dependencyGraph

import (
	"testing"
)

func TestGraph(t *testing.T) {
	graph := NewGraph[string](true)

	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")

	err := graph.AddDependency("A", "B")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = graph.AddDependency("B", "C")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	graph2 := NewGraph[string](false) 
	graph2.AddNode("X")
	graph2.AddNode("Y")
	graph2.AddNode("Z")

	err = graph2.AddDependency("X", "Y")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	graph2.AddDependency("Y", "Z")
	err = graph2.AddDependency("Z", "X")
	if err == nil {
		t.Errorf("Expected circular dependency error, got none")
	}

	err = graph.RemoveNode("C")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = graph.RemoveDependency("A", "B")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}


func TestGraphEntryNodes(t *testing.T) {
	graph := NewGraph[string](true) 

	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")
	graph.AddNode("F")
	graph.AddNode("G")

	graph.AddDependency("A", "B")
	graph.AddDependency("B", "C")
	graph.AddDependency("C", "A")
	graph.AddDependency("D", "A")
	graph.AddDependency("E", "B")
	graph.AddDependency("F", "G")
	graph.RemoveNode("F")

	entryNodes := graph.EntryNodes()

	expectedEntryNodes := []string{"D", "E", "G"}
	for _, node := range expectedEntryNodes {
		found := false
		for _, entry := range entryNodes {
			if entry == node {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected entry node %s not found", node)
		}
	}
}

func TestGraphSplit(t *testing.T) {
	graph := NewGraph[string](true) 
	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddDependency("A", "B")
	graph.AddDependency("A", "C")
	graph.AddDependency("B", "D")
	graph.AddDependency("C", "D")
	graph.RemoveNode("B")
	graph.RemoveNode("C")

	entryNodes := graph.EntryNodes()

	expectedEntryNodes := []string{"A", "D"}
	for _, node := range expectedEntryNodes {
		found := false
		for _, entry := range entryNodes {
			if entry == node {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected entry node %s not found", node)
		}
	}


}

