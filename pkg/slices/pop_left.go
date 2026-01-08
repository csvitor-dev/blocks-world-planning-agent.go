package slices

import "errors"

func PopLeft[T any](slice []T) ([]T, T, error) {
	var value T

	if len(slice) == 0 {
		return slice, value, errors.New("the data structure is empty.")
	}
	return slice[1:], slice[0], nil
}
