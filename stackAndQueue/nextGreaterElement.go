package stackandqueue

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	stack := []int{}

	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			s := stack[len(stack)-1]
			m[s] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	for len(stack) > 0 {
		m[stack[len(stack)-1]] = -1
		stack = stack[:len(stack)-1]
	}

	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = m[nums1[i]]
	}

	return res
}

// https://leetcode.com/problems/next-greater-element-i/
