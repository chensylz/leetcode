package hash

// https://leetcode.cn/problems/two-sum/

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		value, ok := m[target-n]
		if ok {
			return []int{value, i}
		}
		m[n] = i
	}
	return []int{}
}
