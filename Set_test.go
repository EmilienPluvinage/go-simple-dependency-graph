package dependencyGraph

import (
	"testing"
)

func TestSetOperations(t *testing.T) {
    node1 := &node{name: "Node1"}
    node2 := &node{name: "Node2"}
    node3 := &node{name: "Node3"}

    s := newSet(node1, node2)
    if !s.contains(node1) || !s.contains(node2) {
        t.Errorf("newSet failed: expected nodes not present in set")
    }

    s.add(node3)
    if !s.contains(node3) {
        t.Errorf("add failed: added node not present in set")
    }

    s.remove(node2)
    if s.contains(node2) {
        t.Errorf("remove failed: removed node still present in set")
    }

    if !s.contains(node1) {
        t.Errorf("contains failed: expected node1 to be present")
    }
    if s.contains(node2) {
        t.Errorf("contains failed: expected node2 to be removed")
    }

    s2 := newSet(node2, node3)
    s.concat(s2)
    if !s.contains(node2) || !s.contains(node3) {
        t.Errorf("concat failed: expected nodes not present in set after concatenation")
    }

    slice := s.toSlice()
    if len(slice) != len(s) {
        t.Errorf("toSlice failed: slice length mismatch")
    }
    for _, n := range slice {
        if !s.contains(n) {
            t.Errorf("toSlice failed: node %v not present in set", n)
        }
    }
}
