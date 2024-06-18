package linkedlist

// https://leetcode.cn/problems/intersection-of-two-linked-lists-lcci/description/

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeA := headA
	nodeB := headB

	lenA, lenB := 0, 0
	for nodeA != nil {
		lenA++
		nodeA = nodeA.Next
	}
	for nodeB != nil {
		lenB++
		nodeB = nodeB.Next
	}
	var longNode *ListNode
	var shortNode *ListNode
	var size int
	if lenA > lenB {
		longNode = headA
		shortNode = headB
		size = lenA - lenB
	} else {
		longNode = headB
		shortNode = headA
		size = lenB - lenA
	}
	// move to same begin
	for i := 0; i < size; i++ {
		longNode = longNode.Next
	}
	for longNode != shortNode {
		longNode = longNode.Next
		shortNode = shortNode.Next
	}
	return longNode
}
