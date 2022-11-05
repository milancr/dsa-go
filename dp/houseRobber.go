package dp

// recurse
func rob(nums []int) int {
	m := map[int]int{}
	return steal(0, nums, m)
}

func steal(cur int, nums []int, memo map[int]int) int {
	// base case
	if cur >= len(nums) {
		return 0
	}
	// if visited return val
	if _, ok := memo[cur]; ok {
		return memo[cur]
	}

	ans := max(steal(cur+1, nums, memo), steal(cur+2, nums, memo)+nums[cur])

	memo[cur] = ans
	return ans
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func robDp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	amt := make([]int, len(nums))

	for i := len(nums) - 1; i >= 0; i-- {
		amt[i] = max(nums[i+1], nums[i+2]+nums[i])
	}
	return amt[0]
}
