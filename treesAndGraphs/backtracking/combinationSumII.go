package backtracking

// challenge, memoize the problem to improve it, run test cases | build test file, log TC, build benchmark for before and after. Make it faster

func combinationSum2(candidates []int, target int) [][]int {
	results := [][]int{}
	// memo used for dp here in order to store prev values used in recurse to
	//
	memo := map[int]int{}

	for i := range candidates {
		recurse(i, target, []int{}, candidates, results)
	}

	return results
}

func recurse(idx, target int, ans, candidates []int, results [][]int) {
	if target-candidates[idx] == 0 {
		ans = append(ans, candidates[idx])
		results = append(results, append([]int{}, ans...))
		return
	}

	if target-candidates[idx] < 0 {
		return
	}

	target = target - candidates[idx]
	ans = append(ans, candidates[idx])

	for j := idx + 1; j < len(candidates); j++ {
		recurse(j, target, ans, candidates)
	}
}
