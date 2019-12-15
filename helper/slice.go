package helper

// SliceDeduplication去除切片重复元素
func SliceDeduplication(slices []string) []string {
	var result = make([]string, 0)
	var m = make(map[string]bool)
	for _, s := range slices {
		if !m[s] {
			m[s] = true
			result = append(result, s)
		}
	}
	return result
}
