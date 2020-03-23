package leetcode

import (
	toolkit "github.com/byyam/go-leetcode-toolkit"
	"testing"
)

func Test_middleNode(t *testing.T) {
	l1 := &toolkit.ListNode{Val: -1}
	l1.AppendList([]int{1, 2, 3, 4})
	l1.Next.String()
	rst := middleNode(l1.Next)
	if rst == nil {
		t.Fatalf("return nil")
	}
	if rst.Val != 3 {
		t.Fatalf("get middle node failed:%d", rst.Val)
	}
	t.Logf("%+v", rst)

}
