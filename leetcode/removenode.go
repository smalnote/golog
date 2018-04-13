package leetcode

// Given a linked list, remove the n-th node from the end
// of list and return its head.
// Example: 1->2->3->4->5, and n=2 => 1->2->4->5
// Note: Given n will always be valid.
// Follow up: Could you do this in one pass.
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	var preNode *ListNode
	rmNode, cur := head, head
	count := 0
	for cur != nil {
		if count >= n {
			preNode, rmNode = rmNode, rmNode.Next
		}
		cur = cur.Next
		count++
	}
	var newHead *ListNode
	if rmNode == head {
		newHead = rmNode.Next
		rmNode.Next = nil
	} else {
		preNode.Next = rmNode.Next
		newHead = head
	}
	return newHead
}
