package dependencyGraph

type node struct {
	name         interface{}
	dependencies set
	dependants   set
}

func newNode(_name interface{}) *node {
	return &node{
		name:         _name,
		dependencies: set{},
		dependants:   set{},
	}
}

func (n *node) removeNode() {
	for dep, _ := range n.dependants {
		dep.dependencies.remove(n)
	}

	for dep, _ := range n.dependencies {
		dep.dependants.remove(n)
	}
}

func (n *node) dependenciesOf(checkedNodes *set) *set {
	checkedNodes.add(n)
	output := &set{}
	for dep, _ := range n.dependencies {
		output.add(dep)
		if !checkedNodes.contains(dep) {
			temp := dep.dependenciesOf(checkedNodes)
			output.concat(temp)
		}
	}
	return output
}

func (n *node) dependantsOf(checkedNodes *set) *set {
	checkedNodes.add(n)
	output := &set{}
	for dep, _ := range n.dependants {
		output.add(dep)
		if !checkedNodes.contains(dep) {
			temp := dep.dependantsOf(checkedNodes)
			output.concat(temp)
		}
	}
	return output
}

func (n *node) isEntryNode() bool {
	return len(n.dependants) == 0
}
