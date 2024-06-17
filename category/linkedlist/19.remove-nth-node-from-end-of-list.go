package linkedlist

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/

// RemoveNthFromEnd 删除链表的倒数第 N 个结点.
// 快慢指针.
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	fast := dummy
	slow := dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	if slow.Next != nil {
		slow.Next = slow.Next.Next
	}

	return dummy.Next
}
