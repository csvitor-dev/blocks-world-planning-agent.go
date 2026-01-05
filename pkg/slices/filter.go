package slices

func Filter[T any](slice []T, fn func(T, int) bool) []T {
	var result []T

	for i, v := range slice {
		if ok := fn(v, i); ok {
			result = append(result, v)
		}
	}
	return result
}
