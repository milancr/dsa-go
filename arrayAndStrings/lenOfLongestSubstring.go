package arrayandstrings

func lengthOfLongestSubstring(s string) int {
	longest := 0
	p1, p2 := 0, 0
	m := map[byte]bool{}
	for p2 < len(s) {
		c := s[p2]
		p2++
		for p1 < p2 {
			if _, present := m[c]; !present {
				break
			}
			c2 := s[p1]
			p1++
			delete(m, c2)
		}
		m[c] = true
		if longest < p2-p1 {
			longest = p2 - p1
		}

	}
	return longest
}
