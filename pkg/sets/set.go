package sets

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/csvitor-dev/blocks-world-planning-agent.go/pkg/sorted"
)

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](elements ...T) Set[T] {
	set := make(map[T]struct{})

	for _, element := range elements {
		set[element] = struct{}{}
	}
	return set
}

func (set Set[T]) String() string {
	var stringBuilder string

	for element := range set {
		stringBuilder += fmt.Sprintf("%v, ", element)
	}
	var result string

	if len(stringBuilder) > 0 {
		result = stringBuilder[:len(stringBuilder)-2]
	}
	return "{" + result + "}"
}

func SortedString(intSet Set[int]) string {
	clone := make([]int, 0, len(intSet))

	for v := range intSet {
		clone = append(clone, v)
	}
	sorted.InsertionSort(clone)
	str := make([]string, len(clone))

	for i, v := range clone {
		str[i] = strconv.Itoa(v)
	}
	return "[" + strings.Join(str, ",") + "]"
}
