package sets

import (
	"strconv"
	"strings"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	return map[T]struct{}{}
}

func SortedString(intSet Set[int]) string {
	clone := make([]int, 0, len(intSet))

	for v := range intSet {
		clone = append(clone, v)
	}
	str := make([]string, len(clone))

	for i, v := range clone {
		str[i] = strconv.Itoa(v)
	}
	return "[" + strings.Join(str, ",") + "]"
}
