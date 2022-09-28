package hashing

func canConstruct(ransomNote string, magazine string) bool {
	m := map[string]int{}
	for _, v := range magazine {
		if _, ok := m[string(v)]; !ok {
			m[string(v)] = 1
		} else {
			m[string(v)]++
		}
	}

	for _, v := range ransomNote {
		if _, ok := m[string(v)]; ok && m[string(v)] > 0 {
			m[string(v)]--
		} else {
			return false
		}
	}
	return true
}
