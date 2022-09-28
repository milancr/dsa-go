package hashing

func checkIfPangram(sentence string) bool {
	m := map[string]struct{}{}

	for _, val := range sentence {
		_, ok := m[string(val)]
		if ok == false {
			m[string(val)] = struct{}{}
		}
	}
	if len(m) == 26 {
		return true
	}
	return false
}
