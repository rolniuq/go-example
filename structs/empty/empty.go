package empty

/*
	Empty struct:
	- 0 bytes alloc
		var e empty
		type emptyStruct struct{}

	- Same address
		cause all empty alloc use zerobase in alloc.go

	- Stateless
		meaning not contains anything
*/

type Set[K comparable] map[K]struct{}

func (s Set[K]) Add(k K) {
	s[k] = struct{}{}
}

func (s Set[K]) Remove(k K) {
	delete(s, k)
}

func (s Set[K]) Contains(k K) bool {
	_, ok := s[k]
	return ok
}
