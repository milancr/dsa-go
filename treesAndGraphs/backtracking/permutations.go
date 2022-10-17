package backtracking

// nums := []int{1,2,3}
func permutations(nums []int) [][]int {
	res := [][]int{}

	permutation := make([]int, len(nums))
	visit := make(map[int]bool)

	var backtrack func(int)
	backtrack = func(index int) {
		if index == len(nums) {
			copiedPermutation := make([]int, len(nums))
			copy(copiedPermutation, permutation)
			res = append(res, copiedPermutation)
			return
		}

		for _, i := range nums {
			if !visit[i] {
				visit[i] = true
				permutation[index] = i
				backtrack(index + 1)
				visit[i] = false
			}
		}
	}
	backtrack(0)
	return res
}

// https://leetcode.com/problems/permutations/submissions/
