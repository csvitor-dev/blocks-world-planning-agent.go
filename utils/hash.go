package utils

func Hash(str string) uint32 {
	var h uint32 = 2166136261

	for _, c := range str {
		h ^= uint32(c)
		h *= 16777619
	}
	return h
}
