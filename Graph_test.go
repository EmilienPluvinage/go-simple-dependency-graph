package dependencyGraph

import (
	"testing"
)

func TestGraph(t *testing.T) {
	graph := NewGraph(true)

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

	graph2 := NewGraph(false) 
	graph2.AddNode("X")
	graph2.AddNode("Y")
	graph2.AddNode("Z")

	err = graph2.AddDependency("X", "Y")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = graph2.AddDependency("Y", "Z")
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
	graph := NewGraph(true) 

	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	graph.AddNode("E")
	graph.AddNode("F")
	graph.AddNode("G")

	err := graph.AddDependency("A", "B")
	err = graph.AddDependency("B", "C")
	err = graph.AddDependency("C", "A")
	err = graph.AddDependency("D", "A")
	err = graph.AddDependency("E", "B")
	err = graph.AddDependency("F", "G")
	err = graph.RemoveNode("F")

	entryNodes, err := graph.EntryNodes()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

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
	graph := NewGraph(true) 
	graph.AddNode("A")
	graph.AddNode("B")
	graph.AddNode("C")
	graph.AddNode("D")
	err := graph.AddDependency("A", "B")
	err = graph.AddDependency("A", "C")
	err = graph.AddDependency("B", "D")
	err = graph.AddDependency("C", "D")
	err = graph.RemoveNode("B")
	err = graph.RemoveNode("C")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	entryNodes, err := graph.EntryNodes()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

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

