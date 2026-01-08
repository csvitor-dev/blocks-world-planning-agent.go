package sorted

func InsertionSort[T int | float32 | float64](A []T) {
	for i := 1; i < len(A); i++ {
		v := A[i]
		j := i - 1

		for j >= 0 && A[j] > v {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = v
	}
}
