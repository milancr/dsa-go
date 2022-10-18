package backtracking

func combine(n int, k int) [][]int {

	result := make([][]int, 0)

	var comb func(start int, curComb []int)

	comb = func(start int, curComb []int) {

		// Base case
		if len(curComb) == k {

			// deep copy
			dst := make([]int, k)
			copy(dst, curComb)
			result = append(result, dst)
			return
		}

		for i := start; i <= n; i++ {
			curComb = append(curComb, i)
			comb(i+1, curComb)
			curComb = curComb[:len(curComb)-1]
		}
		return
	}

	comb(1, make([]int, 0))
	return result
}

// https://leetcode.com/problems/combinations/
