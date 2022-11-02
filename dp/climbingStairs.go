package dp

// approach recursion wt memoization
// O(n) sc O(n)
func climbStairs(n int) int {
	memo := make([]int, n+1)
	return climb(0, n, memo)
}

func climb(i, n int, memo []int) int {

	if i > n {
		return 0
	}

	if i == n {
		return 1
	}

	if memo[i] > 0 {
		return memo[i]
	}

	memo[i] = climb(i+1, n, memo) + climb(i+2, n, memo)
	return memo[i]
}

// dp approach
func climbStairsDP(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
