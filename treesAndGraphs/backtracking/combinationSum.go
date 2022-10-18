package backtracking

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}

	backtrack(&ans, candidates, target, []int{}, 0)

	return ans
}

// input = [2,3,6,7] target= 7
// need to include index to avoid duplicates, without it you'll get duplicates in different orders
// will get like [2,3,2], [3, 2, 2]... [7] etc. index eliminated duplicates

func backtrack(ans *[][]int, candidates []int, target int, curArr []int, index int) {
	if sum(curArr) == target {
		*ans = append(*ans, append([]int{}, curArr...))
		return
	}

	if sum(curArr) > target {
		return
	}

	for i := index; i < len(candidates); i++ {
		if sum(curArr) < target {
			curArr = append(curArr, candidates[i])
			backtrack(ans, candidates, target, curArr, i)
			curArr = curArr[:len(curArr)-1]
		}
	}
}

func sum(arr []int) int {
	total := 0
	for _, i := range arr {
		total += i
	}
	return total
}

// subtract int added to combo array from target, if sum arr == target append base case
