package genericset

// Implement Set[T comparable] as a map[T]struct{}.
type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return Set[T]{}
}

// Add an element to a Set.
func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

// Check if an element is in a Set.
func (s Set[T]) Contains(elem T) bool {
	_, ok := s[elem]
    return ok
}

// Check if two sets have the same elements.
func (s Set[T]) Equals(other Set[T]) bool {
	if len(s) != len(other) {
        return false
    }
    
	for val, _ := range s {
        if _, contains := other[val]; !contains {
            return false
        }
    }

    return true
}
