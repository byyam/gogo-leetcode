package leetcode

import (
	"github.com/byyam/go-leetcode-toolkit"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &toolkit.ListNode{Val: -1}
	l1.AppendList([]int{5, 1, 3})
	l1.Next.String()
	l2 := &toolkit.ListNode{Val: -1}
	l2.AppendList([]int{5, 8, 6, 1})
	l2.Next.String()

	l0 := addTwoNumbers(l1.Next, l2.Next)
	l0.String()
	if ok := l0.CompareList([]int{0, 0, 0, 2}); !ok {
		t.Fatalf("compare list error!")
	}
	t.Log("ok!")
}
