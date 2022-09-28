package hashing

// nums range is 0 -> len(arr), add
func missingNumber(nums []int) int {
	n := len(nums)
	mn := n * (n + 1) / 2
	for _, v := range nums {
		mn -= v
	}
	return mn
}
