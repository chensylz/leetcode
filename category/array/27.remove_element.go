package array

// 27. 移除元素
// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val 不同的元素的数量。
// 假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：
// 更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
// 返回 k。

// RemoveElement 两层循环法.
// 如果遇到相同的，后面的直接往前移动一位
// 需要注意的是，移动后整体长度会减少一位，那么当前游标的位置则也会退一位
func RemoveElement(nums []int, val int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] == val {
			// 往前移动一位
			for j := i; j < length-1; j++ {
				nums[j] = nums[j+1]
			}
			length--
			i--
		}
	}
	return length
}

// RemoveElementDoublePointer 快慢指针.
// 通过一个快指针和慢指针在一个for循环下完成两个for循环的工作。
// 快指针：寻找新数组的元素，新数组就是不含有目标元素的数组
// 慢指针：指向更新 新数组下标的位置
func RemoveElementDoublePointer(nums []int, val int) int {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
