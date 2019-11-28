package leetcode

import . "github.com/byyam/go-leetcode-toolkit"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l0 := &ListNode{Val: -1}
	p0 := l0
	p1 := l1
	p2 := l2
	x, y, carry := 0, 0, 0
	for p1 != nil || p2 != nil || carry > 0 {
		x = 0
		y = 0
		if p1 != nil {
			x = p1.Val
			p1 = p1.Next
		}
		if p2 != nil {
			y = p2.Val
			p2 = p2.Next
		}
		sum := x + y + carry
		carry = sum / 10
		sum = sum % 10
		p0.Next = &ListNode{Val: sum}
		p0 = p0.Next
	}
	return l0.Next
}
