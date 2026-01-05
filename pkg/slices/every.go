package slices

func Every[T any](slice []T, fn func(T, int) bool) bool {
	result := Filter(slice, fn)

	return len(result) == len(slice)
}
