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

	var listNode ListNode = addTwoListNode(&ListNode{7, nil}, &ListNode{8, nil})

	if (listNode.Val != 15) {
		t.Fatal("fail")
	}
}
