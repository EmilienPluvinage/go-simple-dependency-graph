package dependencyGraph

type set map[*node]struct{}

func newSet(nodes ...*node) set {
	s := make(set)
	for _, n := range nodes {
		s.add(n)
	}
	return s
}

func (s *set) add(elem *node) {
	(*s)[elem] = struct{}{}
}

func (s *set) remove(elem *node) {
	delete(*s, elem)
}

func (s *set) contains(elem *node) bool {
	_, exists := (*s)[elem]
	return exists
}

func (s *set) concat(n *set) {
	for key := range *n {
		s.add(key)
	}
}

func (s *set) toSlice() []*node {
	slice := make([]*node, 0, len(*s))
	for key := range *s {
		slice = append(slice, key)
	}
	return slice
}