package hashing

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		if _, ok := m[target-v]; ok {
			return []int{i, m[target-v]}
		}
		m[v] = i
	}
	return []int{}
}
