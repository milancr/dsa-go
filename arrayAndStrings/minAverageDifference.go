package arrayandstrings

func minimumAverageDifference(nums []int) int {
	// Initialize sum var, left window running sum var, and right average var
	// Sum the array
	sum, left, rightAvg := 0, 0, 0
	for _, val := range nums {
		sum += val
	}

	// Set the minimum value equal to the sum
	// Initialize idx var to store the index val of the running miniumum value between averages
	min := sum
	idx := 0
	// Begin loop at 1 to avoid 0 division
	for i := 1; i <= len(nums); i++ {
		left += nums[i-1]
		right := sum - left

		leftAvg := left / i

		// Avoid 0 division err
		if i == len(nums) {
			rightAvg = 0
		} else {
			rightAvg = right / (len(nums) - i)
		}

		// Calculate absolute difference between averages
		diff := abs(leftAvg - rightAvg)
		// Set new minimum and store index of new minimum
		if diff < min {
			min = diff
			idx = i - 1
		}
	}
	return idx
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
