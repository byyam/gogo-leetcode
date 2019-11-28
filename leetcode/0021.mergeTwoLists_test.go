package leetcode

import (
	toolkit "github.com/byyam/go-leetcode-toolkit"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	l1 := &toolkit.ListNode{Val: -1}
	l1.AppendList([]int{1, 2, 4})
	l1.Next.String()
	l2 := &toolkit.ListNode{Val: -1}
	l2.AppendList([]int{1, 3, 4})
	l2.Next.String()

	l0 := mergeTwoLists(l1.Next, l2.Next)
	l0.String()
	if ok := l0.CompareList([]int{1, 1, 2, 3, 4, 4}); !ok {
		t.Fatalf("compare list error!")
	}
	t.Log("ok!")
}
