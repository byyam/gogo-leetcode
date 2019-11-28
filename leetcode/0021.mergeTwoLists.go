package leetcode

import . "github.com/byyam/go-leetcode-toolkit"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l0 := &ListNode{Val: -1}
	p0 := l0
	p1 := l1
	p2 := l2
	for p1 != nil || p2 != nil {
		if p1 == nil {
			p0.Next = p2
			p2 = p2.Next
		} else if p2 == nil {
			p0.Next = p1
			p1 = p1.Next
		} else {
			if p1.Val < p2.Val {
				p0.Next = p1
				p1 = p1.Next
			} else {
				p0.Next = p2
				p2 = p2.Next
			}
		}
		p0 = p0.Next
	}
	return l0.Next
}
