package slices

func Map[T any, U any](slice []T, fn func(T, int) U) []U {
	result := make([]U, len(slice))

	for i, v := range slice {
		result[i] = fn(v, i)
	}
	return result
}
