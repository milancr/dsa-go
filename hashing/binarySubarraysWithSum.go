package hashing

func numSubarraysWithSum(nums []int, goal int) int {
	m, count, sum := map[int]int{}, 0, 0
	m[0] = 1
	for _, v := range nums {
		sum += v
		if _, ok := m[sum-goal]; ok {
			count += m[sum-goal]
		}
		m[sum]++
	}
	return count
}

// sums occurrences
// m[0] = 1
// m[1] = 1
// m[2] = 1
// m[3] = 1 <- final add
// [1,0,1,0,1]
// sum 1, 1, 2, 2, 3
//                    +1 +2 +1 == 4
// sum - goal: -1, -1, 0, 0, 1

// c: add one to count and one to m[0]
// p : add one to count and 1 to m[1]

// https://leetcode.com/problems/binary-subarrays-with-sum/
