package hashing

// initial impression is to load these into sets and then create a third set for the overlap

func intersection(nums1 []int, nums2 []int) []int {
	s1 := addToSet(nums1)
	s2 := addToSet(nums2)

	var ans []int

	for k := range s1 {
		_, ok := s2[k]
		if ok == true {
			ans = append(ans, k)
		}
	}

	return ans
}

func addToSet(arr []int) map[int]struct{} {
	m := map[int]struct{}{}
	for _, val := range arr {
		_, ok := m[val]
		if ok == false {
			m[val] = struct{}{}
		}
	}
	return m
}

// https://leetcode.com/problems/intersection-of-two-arrays/
