package stackandqueue

func makeGood(s string) string {
	stack := []rune{}
	for k, v := range s {
		if k > 0 && len(stack) > 0 {
			diff := stack[len(stack)-1] - v
			if abs(int(diff)) == 32 {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		stack = append(stack, v)
	}
	str := string(stack)
	return str
}

func abs(n1 int) int {
	if n1 < 0 {
		return n1 * -1
	}
	return n1

}
