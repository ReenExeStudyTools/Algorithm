package algorithm

import (
	"reflect"
	"testing"
)

func TestPlayer(t *testing.T) {
	if !reflect.DeepEqual(findIndexesForSum([]int{1, 2}, 3), []int{0, 1}) {
		t.Fatal("fail")
	}
}
