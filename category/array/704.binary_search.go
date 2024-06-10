package array

// 704. 二分查找
//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
//示例 1:
//	输入: nums = [-1,0,3,5,9,12], target = 9
//	输出: 4
//	解释: 9 出现在 nums 中并且下标为 4
//示例 2:
// 	输入: nums = [-1,0,3,5,9,12], target = 2
//	输出: -1
//	解释: 2 不存在 nums 中因此返回 -1

func Search(nums []int, target int) int {
	return BinarySearch(nums, target, 0, len(nums)-1)
}

func BinarySearch(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return BinarySearch(nums, target, left, mid-1)
	}
	if nums[mid] < target {
		return BinarySearch(nums, target, mid+1, right)
	}
	return -1
}

// 除此之外，二分查找还存在两种变形，
// 01.找到重复元素的第一个
// 02.找到重复元素的最后一个

// SearchFirst 找到重复元素的第一个.
func SearchFirst(nums []int, target int) int {
	return BinarySearchFirst(nums, target, 0, len(nums)-1)
}

func BinarySearchFirst(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if (mid == 0 || nums[mid-1] != target) && nums[mid] == target {
		return mid
	}
	if nums[mid] >= target {
		return BinarySearch(nums, target, left, mid-1)
	}
	if nums[mid] < target {
		return BinarySearch(nums, target, mid+1, right)
	}
	return -1
}

func SearchLast(nums []int, target int) int {
	return BinarySearchLast(nums, target, 0, len(nums)-1)
}

func BinarySearchLast(nums []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if (mid == len(nums)-1 || nums[mid+1] != target) && nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return BinarySearch(nums, target, left, mid-1)
	}
	if nums[mid] <= target {
		return BinarySearch(nums, target, mid+1, right)
	}
	return -1
}
