package dependencyGraph

type node struct {
	name         string
	dependencies []*node
	dependants   []*node
}

func newNode(_name string) *node {
	return &node{
		name:         _name,
		dependencies: []*node{},
		dependants:   []*node{},
	}
}

func (n node) directDependantsOf() set {
	return newSet(n.dependants...)
}

func (n node) directDependenciesOf() set {
	return newSet(n.dependencies...)
}

func (n *node) addDependency(d *node) {
	n.dependencies = append(n.dependencies, d)
}

func (n *node) addDependant(d *node) {
	n.dependants = append(n.dependants, d)
}

func (n *node) removeNode() ([]*node, []*node, error) {
	for _, dep := range n.dependants {
		dep.dependencies = deleteByValue(dep.dependencies, n)
	}

	for _, dep := range n.dependencies {
		dep.dependants = deleteByValue(dep.dependants, n)
	}

	return n.dependencies, n.dependants, nil
}

func (n *node) dependenciesOf(checkedNodes *set) set {
	checkedNodes.add(n)
	output := &set{}
	for _, dep := range n.dependencies {
		output.add(dep)
		if !checkedNodes.contains(dep) {
			temp := dep.dependenciesOf(checkedNodes)
			output.concat(temp)
		}
	}
	return *output
}

func (n *node) dependantsOf(checkedNodes *set) set {
	checkedNodes.add(n)
	output := &set{}
	for _, dep := range n.dependants {
		output.add(dep)
		if !checkedNodes.contains(dep) {
			temp := dep.dependantsOf(checkedNodes)
			output.concat(temp)
		}
	}
	return *output
}

func (n *node) entryNodes(checkedNodes *set, entryNodes *set) {
	checkedNodes.add(n)
	if len(n.dependants) == 0 {
		entryNodes.add(n)
	}

	for _, dep := range n.dependants {
		if !checkedNodes.contains(dep) {
			dep.entryNodes(checkedNodes, entryNodes)
		}
	}

	for _, dep := range n.dependencies {
		if !checkedNodes.contains(dep) {
			dep.entryNodes(checkedNodes, entryNodes)
		}
	}
}

func (n *node) findNode(_name string, checkedNodes *set) (*node, error) {

	if n.name == _name {
		return n, nil
	}

	checkedNodes.add(n)

	for _, dep := range n.dependencies {
		if !checkedNodes.contains(dep) {
			found, err := dep.findNode(_name, checkedNodes)
			if found != nil {
				return found, err
			}
			if err != nil {
				return nil, err
			}
		}
	}

	for _, dep := range n.dependants {
		if !checkedNodes.contains(dep) {
			found, err := dep.findNode(_name, checkedNodes)
			if found != nil {
				return found, err
			}
			if err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}
