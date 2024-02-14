package dependencyGraph

import (
	"errors"
	"reflect"
)

type Graph[T comparable] struct {
	allowCircular bool
	nodeList map[T]*node
}


func NewGraph[T comparable](allowCircular bool) *Graph[T] {
	return &Graph[T]{
		allowCircular,
		make(map[T]*node),
	}
}

func (g * Graph[T]) AddNode(_name T) error {
	if g.HasNode(_name) {
		return errors.New("Node already exists")
	}

	nodeToAdd := newNode(_name)
	g.nodeList[_name] = nodeToAdd
	return nil
}

func (g * Graph[T]) HasNode(_name T) bool {
	return g.nodeList[_name] != nil
}

func (g * Graph[T]) RemoveNode(_name T) error {
	nodeToRemove := g.nodeList[_name]
	if nodeToRemove == nil {
		return errors.New("Node doesn't exist")
	}
	
	nodeToRemove.removeNode()
	delete(g.nodeList, _name)
	return nil
}

func (g * Graph[T]) Size() int {
	return len(g.nodeList)
}

func (g * Graph[T]) EntryNodes() []T {
	entryNodes := []T{}
	for key, value := range g.nodeList {
		if value.isEntryNode() {
			entryNodes = append(entryNodes, key)
		}
	}
	return entryNodes
}

func (g * Graph[T]) DependenciesOf(_name T) ([]T, error) {
	selectedNode := g.nodeList[_name]
	if selectedNode == nil {
		return []T{}, errors.New("Node doesn't exist")
	}
	
	dependencies := selectedNode.dependenciesOf(&set{})
	return nodesTo[T](dependencies.toSlice()), nil
}

func (g * Graph[T]) DependantsOf(_name T) ([]T, error) {
	selectedNode := g.nodeList[_name]
	if selectedNode == nil {
		return []T{}, errors.New("Node doesn't exist")
	}
	
	dependants := selectedNode.dependantsOf(&set{})
	return nodesTo[T](dependants.toSlice()), nil
}

func (g * Graph[T]) DirectDependenciesOf(_name T) ([]T, error) {
	selectedNode := g.nodeList[_name]
	if selectedNode == nil {
		return []T{}, errors.New("Node doesn't exist")
	}
	dependencies := make([]T, 0, len(selectedNode.dependencies))
	for key := range selectedNode.dependencies {
		dependencies = append(dependencies, key.name.(T))
	}
	return dependencies, nil
}

func (g * Graph[T]) DirectDependantsOf(_name T) ([]T, error) {
	selectedNode := g.nodeList[_name]
	if selectedNode == nil {
		return []T{}, errors.New("Node doesn't exist")
	}
	dependants := make([]T, 0, len(selectedNode.dependants))
	for key := range selectedNode.dependants {
		dependants = append(dependants, key.name.(T))
	}
	return dependants, nil
}

func (g * Graph[T]) RemoveDependency(from T, to T) error {
	fromNode := g.nodeList[from]
	if fromNode == nil {
		return errors.New("Node doesn't exist")
	}

	var toNode *node
	for key := range fromNode.dependencies {
		if reflect.DeepEqual(key.name.(T), to) {
			toNode = key
			break
		}
	}

	if toNode == nil {
		return errors.New("Unknown dependency")
	}
	
	fromNode.dependencies.remove(toNode)
	toNode.dependants.remove(fromNode)

	return nil
}

func (g * Graph[T]) AddDependency(from T, to T) error {
	if reflect.DeepEqual(from, to) {
		return errors.New("Circular dependency")
	}

	fromNode := g.nodeList[from]
	if fromNode == nil {
		return errors.New("Node doesn't exist")
	}

	toNode := g.nodeList[to]
	if toNode == nil {
		return errors.New("Node doesn't exist")
	}


	directDependencies := fromNode.dependencies
	if directDependencies.contains(toNode) {
		return errors.New("Already a dependency")
	}

	if !g.allowCircular {
		allDependants := fromNode.dependantsOf(&set{})
		if allDependants.contains(toNode) {
			return errors.New("Circular dependency")
		}
	}
	
	fromNode.dependencies.add(toNode)
	toNode.dependants.add(fromNode)

	return nil
}