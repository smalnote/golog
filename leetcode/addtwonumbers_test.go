package leetcode

import (
	"fmt"
	"testing"
)

// You are given two non-empty linked lists representing
// tow non-negative integers.
// The digits are store in reverse order and each of their
// nodes contain a single digit. Add the tow numbers and
// return it as a linked list.
// You may assume the two numbers do not contain any leading
// zero, except the number 0 itself.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	e1 := l1
	e2 := l2
	carryover := 0
	zero := &ListNode{
		0,
		nil,
	}
	var res, last *ListNode
	for e1 != zero || e2 != zero {
		v := e1.Val + e2.Val + carryover
		carryover = v / 10
		v = v % 10
		n := &ListNode{
			v,
			nil,
		}
		if last == nil {
			res, last = n, n
		} else {
			last.Next, last = n, n
		}
		if e1.Next == nil {
			e1 = zero
		} else {
			e1 = e1.Next
		}
		if e2.Next == nil {
			e2 = zero
		} else {
			e2 = e2.Next
		}
	}
	if carryover > 0 {
		last.Next = &ListNode{
			carryover,
			nil,
		}
		last = last.Next
	}
	return res
}

func (l *ListNode) String() string {
	var s string
	for e := l; e != nil; e = e.Next {
		if e == l {
			s = fmt.Sprintf("%d", e.Val)
		} else {
			s = fmt.Sprintf("%s->%d", s, e.Val)
		}
	}
	return s
}

func TestAddTwoNumbers(t *testing.T) {
	testcases := []struct {
		l1  *ListNode
		l2  *ListNode
		res *ListNode
	}{
		{
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			&ListNode{0, nil},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			&ListNode{5, &ListNode{6, &ListNode{4, nil}}},
		},
		{
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			&ListNode{0, nil},
			&ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		},
		{
			&ListNode{5, &ListNode{2, nil}},
			&ListNode{5, &ListNode{7, nil}},
			&ListNode{0, &ListNode{0, &ListNode{1, nil}}},
		},
	}

	for i, c := range testcases {
		got := addTwoNumbers(c.l1, c.l2).String()
		expected := c.res.String()
		if expected != got {
			t.Errorf("#%d %s + %s expected %s, got %s ", i, c.l1, c.l2, expected, got)
		} else {
			t.Log(got)
		}
	}
}
