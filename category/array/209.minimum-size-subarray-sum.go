package array

import "math"

// https://leetcode.cn/problems/minimum-size-subarray-sum/description/

func MinSubArrayLen(target int, nums []int) int {
	minSize := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		size := 0
		total := 0
		for j := i; j < len(nums); j++ {
			total += nums[j]
			size++
			if total >= target {
				minSize = min(minSize, size)
				break
			}
		}
	}
	if minSize == math.MaxInt32 {
		return 0
	}
	return minSize
}

// MinSubArrayLenWindow 滑动窗口。
// 不断的调节子序列的起始位置和终止位置，从而得出我们要想的结果
// 在上面解法中，是一个for循环滑动窗口的起始位置，一个for循环为滑动窗口的终止位置，用两个for循环 完成了一个不断搜索区间的过程
// 根据当前子序列和大小的情况，不断调节子序列的起始位置。从而将O(n^2)暴力解法降为O(n)
func MinSubArrayLenWindow(target int, nums []int) int {
	left := 0
	total := 0
	result := len(nums) + 1
	for right := 0; right < len(nums); right++ {
		total += nums[right]
		for total >= target {
			subLen := right - left + 1
			result = min(subLen, result)
			total -= nums[left]
			left++
		}
	}
	if result == len(nums)+1 {
		return 0
	}
	return result
}
