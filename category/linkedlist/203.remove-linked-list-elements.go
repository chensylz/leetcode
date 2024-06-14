package linkedlist

// RemoveElements 移除链表元素.
// https://leetcode.cn/problems/remove-linked-list-elements/
func RemoveElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}
