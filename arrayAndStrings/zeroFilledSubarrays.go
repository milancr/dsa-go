package arrayandstrings

// [0,0,0,0] == [0]: 4 [0,0]: 3 [0,0,0]: 2 [0,0,0,0]: 1
// [0,0,0] == [0]: 3 [0,0]: 2 [0,0,0]: 1
// sliding window, use dp to store subarray lengths in array, reset window length. Upon sliding window completion apply n(n+1)/2 to each subarray length stored in dp and add to sum

func zeroFilledSubarray(nums []int) int64 {
	var dp []int
	length := 0
	for _, val := range nums {
		if val == 0 {
			length++
		} else if length > 0 {
			dp = append(dp, length)
			length = 0
		}
	}
	dp = append(dp, length)
	sum := 0

	for _, val := range dp {
		sum = sum + ((val * (val + 1)) / 2)
	}

	return int64(sum)
}

// https://leetcode.com/problems/number-of-zero-filled-subarrays/
