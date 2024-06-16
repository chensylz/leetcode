package linkedlist

// https://leetcode.cn/problems/swap-nodes-in-pairs/

func SwapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	// 1 -> 2 -> 3 -> 4
	// n1   n2   n3   n4
	for cur.Next != nil && cur.Next.Next != nil {
		tmp := cur.Next            // n1
		tmp1 := cur.Next.Next.Next // n3

		cur.Next = cur.Next.Next  //  cur.Next = n2
		cur.Next.Next = tmp       //  n2(cur.Next).Next = n1
		cur.Next.Next.Next = tmp1 //  n1.Next = n3

		cur = cur.Next.Next //  cur = n1
	}

	return dummy.Next
}
