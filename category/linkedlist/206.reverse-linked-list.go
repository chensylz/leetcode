package linkedlist

// https://leetcode.cn/problems/reverse-linked-list/description/

func ReverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
