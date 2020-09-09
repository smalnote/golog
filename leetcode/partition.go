/** given a linked list and a value x, partition it such
  * tha all nodes less than x come before nodes greater
  * than or equal to x
  * you should preserve the original relative order of the nodes 
  * in each of the two partions.
  */
package leetcode

func partition(head *ListNode, x int) *ListNode {
	n := &ListNode{0, nil}
	m := &ListNode{0, nil}
	ntail := n
	mtail := m
	for k := head; k != nil; k = k.Next {
		if k.Val < x {
			ntail.Next = k	
			ntail = ntail.Next
		} else {
			mtail.Next = k
			mtail= mtail.Next
		}
	}
	if ntail != nil {
		ntail.Next = nil
	}
	if mtail != nil {
		mtail.Next = nil
	}
	var newHead *ListNode = nil
	if n.Next != nil {
		newHead = n.Next
		ntail.Next = m.Next
	} else if m.Next != nil {
		newHead = m.Next 
	}
	return newHead
}