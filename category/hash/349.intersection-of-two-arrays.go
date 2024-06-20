package hash

// https://leetcode.cn/problems/intersection-of-two-arrays/

func Intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] += 1
	}
	result := make([]int, 0)
	m2 := make(map[int]struct{})
	for _, v := range nums2 {
		_, ok1 := m2[v]
		if ok1 {
			continue
		}
		_, ok := m[v]
		if ok {
			result = append(result, v)
			m2[v] = struct{}{}
		}
	}
	return result
}
