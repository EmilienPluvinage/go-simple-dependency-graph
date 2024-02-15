package dependencyGraph

func deleteByValue(arr []*node, value *node) []*node {
	index := -1
	for i, elem := range arr {
		if elem == value {
			index = i
			break
		}
	}

	if index != -1 {
		result := make([]*node, 0, len(arr)-1)
		result = append(result, arr[:index]...)
		result = append(result, arr[index+1:]...)

		return result
	}

	return arr
}

func deleteByIndex(arr []*node, index int) []*node {

	if index < 0 || index >= len(arr) {
		return arr
	}

	result := make([]*node, 0, len(arr)-1)
	result = append(result, arr[:index]...)
	result = append(result, arr[index+1:]...)

	return result
}

func nodesTo[T any](nodes []*node) []T {
	result := make([]T, len(nodes))
	for i, n := range nodes {
		result[i] = n.name.(T)
	}
	return result
}
