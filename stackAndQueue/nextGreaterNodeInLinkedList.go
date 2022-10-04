package stackandqueue

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	// Stack for our operations, we'll keep runs of lower elements here.
	stack := []int{}
	// vals we'll store the vals we take from the ll.
	vals := []int{}
	// Res will be our result array (up to 10k per the question).
	res := [10000]int{}
	// Track the idx we're working through.
	idx := 0
	cur := head
	// While we have vals in our ll.
	for cur != nil {
		// We check, if we have a stack and the current val > whats on the top of the stack.
		for len(stack) > 0 && cur.Val > vals[stack[len(stack)-1]] {
			// We pop the idx from the stack, assign the idx in our res to = the current val.
			res[stack[len(stack)-1]] = cur.Val
			stack = stack[:len(stack)-1]
		}
		// Append the cur idx to the stack (we use idxs not values for ease in assigning res + to handle duplicates!).
		stack = append(stack, idx)
		// Updates idx, append val and move to next node.
		idx++
		vals = append(vals, cur.Val)
		cur = cur.Next
	}
	// Return the res for the number of vals we saw
	return res[:idx]
}

// https://leetcode.com/problems/next-greater-node-in-linked-list/
