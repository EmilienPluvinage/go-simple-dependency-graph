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

func (g * Graph[T]) EntryNodes() ([]T, error) {
	entryNodes := &set{}
	for _, value := range g.nodeList {
		checkedNodes := &set{}
		value.entryNodes(checkedNodes, entryNodes)
	}
	return nodesTo[T](entryNodes.toSlice()), nil
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

	result := make([]T, 0, len(selectedNode.dependencies)-1)
	for index, dep := range selectedNode.dependencies {
		result[index] = dep.name.(T)
	}
	return result, nil
}

func (g * Graph[T]) DirectDependantsOf(_name T) ([]T, error) {
	selectedNode := g.nodeList[_name]
	if selectedNode == nil {
		return []T{}, errors.New("Node doesn't exist")
	}

	result := make([]T, 0, len(selectedNode.dependants)-1)
	for index, dep := range selectedNode.dependants {
		result[index] = dep.name.(T)
	}
	return result, nil
}

func (g * Graph[T]) RemoveDependency(from T, to T) error {
	fromNode := g.nodeList[from]
	if fromNode == nil {
		return errors.New("Node doesn't exist")
	}

	var toNode *node
	for _, value := range fromNode.dependencies {
		if reflect.DeepEqual(value.name.(T), to) {
			toNode = value
			break
		}
	}

	if toNode == nil {
		return errors.New("Unknown dependency")
	}
	
	fromNode.dependencies = deleteByValue(fromNode.dependencies, toNode)
	toNode.dependants = deleteByValue(toNode.dependants, fromNode)

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


	directDependencies := fromNode.directDependenciesOf()
	if directDependencies.contains(toNode) {
		return errors.New("Already a dependency")
	}

	if !g.allowCircular {
		allDependants := fromNode.dependantsOf(&set{})
		if allDependants.contains(toNode) {
			return errors.New("Circular dependency")
		}
	}
	
	fromNode.addDependency(toNode)
	toNode.addDependant(fromNode)

	return nil
}