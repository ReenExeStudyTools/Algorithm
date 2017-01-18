package algorithm

import (
	"reflect"
	"testing"
)

func TestPlayer(t *testing.T) {
	if !reflect.DeepEqual(findIndexesForSum([]int{1, 2}, 3), []int{0, 1}) {
		t.Fatal("fail")
	}

	if !reflect.DeepEqual(findIndexesForSum([]int{3, 2, 4}, 6), []int{1, 2}) {
		t.Fatal("fail")
	}

	listNode := addTwoListNode(&ListNode{3, nil}, &ListNode{2, nil})

	if (listNode.Val != 5) {
		t.Fatal("fail")
	}

	listNodeNested := addTwoListNode(
		&ListNode{2, &ListNode{8, &ListNode{3, nil}}},
		&ListNode{3, &ListNode{9, &ListNode{4, nil}}},
	)

	if (listNodeNested.Val != 5 || listNodeNested.Next.Val != 7 || listNodeNested.Next.Next.Val != 8) {
		t.Fatal("fail")
	}
}
