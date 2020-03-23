package leetcode

import (
	. "github.com/byyam/go-leetcode-toolkit"
	"log"
)

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	count := 0
	for p := head; p != nil; p = p.Next {
		count++
	}
	index := count/2 + 1
	log.Printf("count=%d, index=%d", count, index)
	p := head
	for i := 1; i < index; i++ {
		p = p.Next
	}
	return p
}
