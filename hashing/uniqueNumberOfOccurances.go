package hashing

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	for _, val := range arr {
		if _, ok := m[val]; ok {
			m[val]++
		} else {
			m[val] = 1
		}
	}

	m2 := map[int]int{}
	for _, v := range m {
		if _, ok := m2[v]; ok {
			return false
		}
		m2[v] = 1
	}
	return true
}
