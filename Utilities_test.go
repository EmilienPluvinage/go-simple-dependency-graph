package dependencyGraph

import (
	"reflect"
	"testing"
)

func TestDeleteByValue(t *testing.T) {
    node1 := &node{name: "Node1"}
    node2 := &node{name: "Node2"}
    node3 := &node{name: "Node3"}
    nodes := []*node{node1, node2, node3}

    result := deleteByValue(nodes, node2)
    expected := []*node{node1, node3}
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("deleteByValue failed: got %v, expected %v", result, expected)
    }

    result = deleteByValue(nodes, &node{name: "NonExistent"})
    if !reflect.DeepEqual(result, nodes) {
        t.Errorf("deleteByValue failed: got %v, expected %v", result, nodes)
    }
}

func TestDeleteByIndex(t *testing.T) {
    node1 := &node{name: "Node1"}
    node2 := &node{name: "Node2"}
    node3 := &node{name: "Node3"}
    nodes := []*node{node1, node2, node3}

    result := deleteByIndex(nodes, 1)
    expected := []*node{node1, node3}
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("deleteByIndex failed: got %v, expected %v", result, expected)
    }

    result = deleteByIndex(nodes, 3)
    if !reflect.DeepEqual(result, nodes) {
        t.Errorf("deleteByIndex failed: got %v, expected %v", result, nodes)
    }
}

func TestNodesToString(t *testing.T) {
    node1 := &node{name: "Node1"}
    node2 := &node{name: "Node2"}
    node3 := &node{name: "Node3"}
    nodes := []*node{node1, node2, node3}

    result := nodesTo[string](nodes)
    expected := []string{"Node1", "Node2", "Node3"}
    if !reflect.DeepEqual(result, expected) {
        t.Errorf("nodesToString failed: got %v, expected %v", result, expected)
    }
}
