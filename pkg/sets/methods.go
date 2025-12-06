package sets

func (set Set[T]) Clone() Set[T] {
	clone := make(Set[T])

	for elem := range set {
		clone[elem] = struct{}{}
	}
	return clone
}

func (set Set[T]) Union(other Set[T]) Set[T] {
	return set
}

func (set Set[T]) Intersect(other Set[T]) Set[T] {
	return set
}

func (set Set[T]) Difference(other Set[T]) Set[T] {
	return set
}

func (set Set[T]) Has(element T) bool {
	return false
}

func (set Set[T]) IsSubsetOf(other Set[T]) bool {
	return false
}

func (set Set[T]) AddFrom(elements []T) {
	for _, elem := range elements {
		set[elem] = struct{}{}
	}
}
