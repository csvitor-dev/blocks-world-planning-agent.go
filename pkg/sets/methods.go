package sets

func (set Set[T]) Clone() Set[T] {
	clone := make(Set[T])

	for elem := range set {
		clone[elem] = struct{}{}
	}
	return clone
}

func (set Set[T]) Union(other Set[T]) Set[T] {
	out := set.Clone()

	for v := range other {
		out[v] = struct{}{}
	}
	return out
}

func (set Set[T]) Intersect(other Set[T]) Set[T] {
	out := make(Set[T])

	for v := range set {
		if _, exists := other[v]; exists {
			out[v] = struct{}{}
		}
	}
	return out
}

func (set Set[T]) Difference(other Set[T]) Set[T] {
	out := make(Set[T])

	for v := range set {
		if _, exists := other[v]; !exists {
			out[v] = struct{}{}
		}
	}
	return out
}

func (set Set[T]) Has(element T) bool {
	_, exists := set[element]
	return exists
}

func (set Set[T]) IsSubsetOf(other Set[T]) bool {
	result := true

	for v := range other {
		if _, exists := set[v]; !exists {
			result = false
			break
		}
	}
	return result
}

func (set Set[T]) Equals(other Set[T]) bool {
	result := true

	for v := range other {
		if _, exists := set[v]; !exists {
			result = false
			break
		}
	}
	return result
}

func (set Set[T]) Add(element T) {
	set[element] = struct{}{}
}

func (set Set[T]) AddFrom(elements []T) {
	for _, element := range elements {
		set.Add(element)
	}
}
