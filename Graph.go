package dependencyGraph

import (
	"errors"
	"reflect"
)

type Graph[T any] struct {
	allowCircular bool
	nodeList []*node
	size int
}


func NewGraph[T any](allowCircular bool) *Graph[T] {
	return &Graph[T]{
		allowCircular,
		[]*node{},
		0,
	}
}

func (g * Graph[T]) findNode(_name T) (*node, int, error) {
	for index, value := range g.nodeList {
		checkedNodes := &set{}
		found, _ := value.findNode(_name, checkedNodes)
		if found != nil {
			return found, index, nil
		}
	}

	return nil, 0, errors.New("Unknown node")
}

func (g * Graph[T]) AddNode(_name T) error {
	hasNode, _ := g.HasNode(_name)
	if hasNode {
		return errors.New("Node already exists")
	}

	nodeToAdd := newNode(_name)

	g.size++
	g.nodeList = append(g.nodeList, nodeToAdd)
	return nil
}

func (g * Graph[T]) HasNode(_name T) (bool, error) {
	found, _, err := g.findNode(_name)
	if err != nil {
		return false, nil
	}
	
	return found != nil, nil
}

func (g * Graph[T]) RemoveNode(_name T) error {
	found, index, err := g.findNode(_name)
	if err != nil {
		return err
	}
	
	dependencies, dependants, err := found.removeNode()
	if err != nil {
		return err
	}

	for _, dep := range dependencies {
		foundDep, _ := g.nodeList[index].findNode(dep.name, &set{})
		if foundDep == nil {
			g.nodeList = append(g.nodeList, dep)
		}
	}

	for _, dep := range dependants {
		foundDep, _ := g.nodeList[index].findNode(dep.name, &set{})
		if foundDep == nil {
			g.nodeList = append(g.nodeList, dep)
		}
	}

	g.size--
	return nil
}

func (g * Graph[T]) Size() (int, error) {
	return g.size, nil
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
	found, _, err := g.findNode(_name)
	if err != nil {
		return nil, err
	}
	
	dependencies := found.dependenciesOf(&set{})
	return nodesTo[T](dependencies.toSlice()), nil
}

func (g * Graph[T]) DependantsOf(_name T) ([]T, error) {
	found, _, err := g.findNode(_name)
	if err != nil {
		return nil, err
	}
	

	dependants := found.dependantsOf(&set{})
	return nodesTo[T](dependants.toSlice()), nil
}

func (g * Graph[T]) DirectDependenciesOf(_name T) ([]T, error) {
	found, _, err := g.findNode(_name)
	if err != nil {
		return []T{}, err
	}
	result := make([]T, 0, len(found.dependencies)-1)
	for index, dep := range found.dependencies {
		result[index] = dep.name.(T)
	}
	return result, nil
}

func (g * Graph[T]) DirectDependantsOf(_name T) ([]T, error) {
	found, _, err := g.findNode(_name)
	if err != nil {
		return []T{}, err
	}
	result := make([]T, 0, len(found.dependants)-1)
	for index, dep := range found.dependants {
		result[index] = dep.name.(T)
	}
	return result, nil
}

func (g * Graph[T]) RemoveDependency(from T, to T) error {
	fromNode, index, err := g.findNode(from)
	if err != nil {
		return err
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

	rootNode := g.nodeList[index]

	AllDependencies := rootNode.dependenciesOf(&set{})
	AllDependants := rootNode.dependantsOf(&set{})
	AllDependants.concat(AllDependencies)
	var foundFrom, foundTo bool

	for key := range AllDependants {
		if key == fromNode {
			foundFrom = true
		}
		if key == toNode {
			foundTo = true
		}
	}

	if !foundFrom {
		g.nodeList = append(g.nodeList, fromNode)
	}

	if !foundTo {
		g.nodeList = append(g.nodeList, toNode)
	}

	return nil
}

func (g * Graph[T]) AddDependency(from T, to T) error {
	if reflect.DeepEqual(from, to) {
		return errors.New("Circular dependency")
	}

	fromNode, fromIndex, err := g.findNode(from)
	if err != nil {
		return err
	}
	toNode, toIndex, err := g.findNode(to)
	if err != nil {
		return err
	}

	if toIndex == fromIndex {
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
	}

	fromNode.addDependency(toNode)
	toNode.addDependant(fromNode)

	if toIndex != fromIndex {
		g.nodeList = deleteByIndex(g.nodeList, toIndex)
	}

	return nil
}