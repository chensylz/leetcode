package array

// https://leetcode.cn/problems/squares-of-a-sorted-array/description/

// SortedSquares 有序数组的平方.
// 时间复杂度: O(n)
// 从两端开始找，如果是绝对值大的，直接放在后面
func SortedSquares(nums []int) []int {
	newArr := make([]int, len(nums))
	left := 0
	right := len(nums) - 1
	index := len(nums) - 1
	for left <= right {
		if abs(nums[right]) > abs(nums[left]) {
			newArr[index] = nums[right] * nums[right]
			right--
		} else {
			newArr[index] = nums[left] * nums[left]
			left++
		}
		index--
	}
	return newArr
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
