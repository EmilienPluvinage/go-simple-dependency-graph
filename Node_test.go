package dependencyGraph

import (
	"testing"
)

func TestNodeOperations(t *testing.T) {
    node1 := newNode("Node1")
    node2 := newNode("Node2")
    node3 := newNode("Node3")

    node1.addDependency(node2)
    node1.addDependency(node3)
    dependencies := node1.directDependenciesOf()
    if !dependencies.contains(node2) || !dependencies.contains(node3) {
        t.Errorf("addDependency or directDependenciesOf failed: expected dependencies not present")
    }

    node2.addDependant(node1)
    node3.addDependant(node1)
	node2dependants := node2.directDependantsOf()
	node3dependants := node3.directDependantsOf()
    if !node2dependants.contains(node1) || !node3dependants.contains(node1) {
        t.Errorf("addDependant or directDependantsOf failed: expected dependants not present")
    }

    node1.removeNode()
	node2dependants = node2.directDependantsOf()
	node3dependants = node3.directDependantsOf()
	if node2dependants.contains(node1) || node3dependants.contains(node1) {
        t.Errorf("removeNode failed: dependencies or dependants not correctly updated")
    }

    dependencies = node1.dependenciesOf(&set{})
    if !dependencies.contains(node2) || !dependencies.contains(node3) {
        t.Errorf("dependenciesOf failed: expected dependencies not present")
    }

	node4 := newNode("Node4")
    node5 := newNode("Node5")
    node6 := newNode("Node6")
	node4.addDependant(node5)
	node5.addDependant(node6)
    node4dependants := node4.dependantsOf(&set{})
    if !node4dependants.contains(node5) || !node4dependants.contains(node6) {
        t.Errorf("dependantsOf failed: expected dependants not present")
    }
	if len(node4dependants.toSlice()) != 2 {
        t.Errorf("dependantsOf failed: incorrect number of dependants")
    }

    entryNodes := newSet()
    node1.entryNodes(&set{}, &entryNodes)
    if !entryNodes.contains(node1) {
        t.Errorf("entryNodes failed: expected entry node not present")
    }

    found, err := node1.findNode("Node2", &set{})
    if found == nil || found != node2 || err != nil {
        t.Errorf("findNode failed: expected node2 to be found without error")
    }
}
