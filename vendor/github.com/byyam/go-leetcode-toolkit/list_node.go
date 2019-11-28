package toolkit

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() {
	p := l
	for p != nil {
		fmt.Printf("->%d", p.Val)
		p = p.Next
	}
	fmt.Println()
}

func (l *ListNode) AppendList(array []int) {
	p := l
	for p.Next != nil {
		p = p.Next
	}
	for _, value := range array {
		p.Next = &ListNode{Val: value}
		p = p.Next
	}
}

func (l *ListNode) CompareList(array []int) bool {
	p := l
	for _, value := range array {
		if p == nil {
			return false
		}
		if p.Val != value {
			return false
		}
		p = p.Next
	}
	if p == nil {
		return true
	}
	return false
}
