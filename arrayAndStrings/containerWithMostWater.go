package arrayandstrings

// two pointer approach keeping track of max water held, the shorter line must be used to calculate how much water can be held. Formula is distance between container edges x min height

func maxArea(height []int) int {
	var i, th, maxh int
	j := len(height) - 1

	for i < j {

		if height[i] < height[j] {
			th = (j - i) * height[i]
			i++
		} else {
			th = (j - i) * height[j]
			j--
		}

		if th > maxh {
			maxh = th
		}
	}

	return maxh
}

// https://leetcode.com/problems/container-with-most-water/
