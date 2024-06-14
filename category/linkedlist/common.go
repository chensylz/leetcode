package linkedlist

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val: val,
	}
}

func NewListWithArray(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := NewListNode(nums[0])
	cur := head
	for i := 1; i < len(nums); i++ {
		cur.Next = NewListNode(nums[i])
		cur = cur.Next
	}
	return head
}

func ExtendLinkedToArrayStr(node *ListNode) string {
	array := make([]string, 0)
	for node != nil {
		array = append(array, strconv.Itoa(node.Val))
		node = node.Next
	}
	return strings.Join(array, ",")
}

func ExtendArrayToStr(nums []int) string {
	arr := make([]string, len(nums))
	for i := range nums {
		arr[i] = strconv.Itoa(nums[i])
	}
	return strings.Join(arr, ",")
}
