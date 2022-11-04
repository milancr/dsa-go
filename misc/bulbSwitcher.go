package misc

import "math"

// There are n bulbs that are initially off. You first turn on all the bulbs, then you turn off every second bulb.
// math based question involving square roots
func bulbSwitcher(n int) int {
	// floors square rooted n,
	// Since squares have an odd number of non-repeating divisors, they will toggle an odd number of times and end up lit up.
	return int(math.Sqrt(float64(n)))
}

// original answer that works but times out. Could definitely be improved
// but it works until n = 10^9
func bulbSwitch(n int) int {
	bulbs := make([]int, n)
	for i := range bulbs {
		bulbs[i] = 1
	}

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	for i := 2; i <= n; i++ {
		j := i - 1
		for j < n {
			bulbs[j] = flip(bulbs[j])
			j += i
		}
	}

	return sum(bulbs)
}

func flip(n int) int {
	if n == 0 {
		return 1
	}
	return 0
}

func sum(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
