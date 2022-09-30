package arrayandstrings

func longestOnes(nums []int, k int) int {
	left, right := 0, 0
	for idx := range nums {
		right = idx
		k -= 1 - nums[right]
		if k < 0 {
			k += 1 - nums[left]
			left++
		}
	}
	return right - left + 1

}

// https://leetcode.com/problems/max-consecutive-ones-iii/
