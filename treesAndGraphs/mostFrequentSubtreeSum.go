package treesandgraphs

// find most frequent sum of subtree, create map to count the freq ( k = sum, v = freq)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	m := map[int]int{}
	// handle nil input
	if root == nil {
		return []int{}
	}

	helper(root, m)
	return getFreq(m)
}

// dfs to return sum and update map
func helper(node *TreeNode, m map[int]int) int {
	sum := node.Val

	if node.Left != nil {
		sum += helper(node.Left, m)
	}

	if node.Right != nil {
		sum += helper(node.Right, m)
	}

	m[sum]++
	return sum
}

// find most frequent sums
func getFreq(m map[int]int) []int {
	ans := []int{}
	max := 0
	for k, v := range m {
		if v > max {
			max = v
			ans = nil
			ans = append(ans, k)
		} else if max == v {
			ans = append(ans, k)
		}
	}
	return ans
}

// https://leetcode.com/problems/most-frequent-subtree-sum/
