package algorithm

import (
	"reflect"
	"testing"
)

func TestPlayer(t *testing.T) {
	expect := []int{1}
	actual := findIndexesForSum([]int{1}, 1)
	if !reflect.DeepEqual(actual, expect) {
		t.Fatal("fail")
	}
}
