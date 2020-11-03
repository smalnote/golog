package leetcode

// 把N上有序链表合并成一个, 把链表头构造成一个优先级队列, 每次从队列取, 取完之后将取的链表的下一个元素入队
import (
	"container/heap"
	"testing"
)

type PQ []*ListNode

func (pq PQ) Len() int { return len(pq) }
func (pq PQ) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}
func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}
func (pq *PQ) Push(elem interface{}) {
	*pq = append(*pq, elem.(*ListNode))
}

func (pq *PQ) Pop() interface{} {
	l := len(*pq) - 1
	elem := (*pq)[l]
	*pq = (*pq)[:l]
	return elem
}

func mergeKLists(lists []*ListNode) *ListNode {
	p := PQ(lists)
	pl := &p
	heap.Init(pl)
	for _, list := range lists {
		if list == nil {
			continue
		}
		heap.Push(pl, list)
	}

	head := &ListNode{0, nil}
	curr := head
	for pl.Len() > 0 {
		elem := heap.Pop(pl).(*ListNode)
		curr.Next = elem
		curr = curr.Next
		if elem.Next != nil {
			heap.Push(pl, elem.Next)
		}
	}

	return head.Next
}

func TestMergeKLists(t *testing.T) {

}
