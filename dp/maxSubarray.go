package dp

func maxSubArray(nums []int) int {

	curArr := nums[0]
	MaxArr := nums[0]

	for i := 1; i < len(nums); i++ {
		curArr = max(nums[i], curArr+nums[i])
		MaxArr = max(MaxArr, curArr)
	}
	return MaxArr
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
