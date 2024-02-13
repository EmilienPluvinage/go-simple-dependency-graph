# Dependency Graph

This is a simple dependency graph package in Go inspired from a similar package in NodeJS : https://github.com/jriecken/dependency-graph

## Installation
To be completed

## Documentation

- NewGraph(allowCircular bool) - instantiates a new graph. allowCircular definies whether this graph allows circular references or not.
- AddNode(name) - adds a node to a graph (without any dependency)
- HasNode(name) - checks if a node exists in the graph
- Size() - return the number of nodes in the graph
- RemoveNode(name) - removes a node and its dependencies
- AddDependency(from, to) - add a dependency between two nodes (will throw an Error if one of the nodes does not exist, if the dependency already exists or if the addition results in a circular reference and allowCircular was set to false)
- RemoveDependency(from, to) - remove a dependency between two nodes
- EntryNodes() - array of nodes that have no dependants (i.e. nothing depends on them).
- DependenciesOf(name) - get an array containing the nodes that the specified node depends on (transitively).
- DependantsOf(name) - get an array containing the nodes that depend on the specified node (transitively).
- DirectDependenciesOf(name) - get an array containing the direct dependencies of the specified node
- DirectDependantsOf(name) (aliased as directDependentsOf) - get an array containing the nodes that directly depend on the specified node

## Example
```
graph := NewGraph(true) 

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

entryNodes, err := graph.EntryNodes()

dependenciesOfA, err := graph.DependenciesOf("A")
```

## Future developments
- Generic node types (instead of just strings)
- Concurrency to make it faster
