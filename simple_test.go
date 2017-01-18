package algorithm

import (
	"reflect"
	"testing"
	"fmt"
)

func TestPlayer(t *testing.T) {
	if !reflect.DeepEqual(findIndexesForSum([]int{1, 2}, 3), []int{0, 1}) {
		t.Fatal("fail")
	}

	fmt.Print(findIndexesForSum([]int{3, 2, 4}, 6))
	if !reflect.DeepEqual(findIndexesForSum([]int{3, 2, 4}, 6), []int{1, 2}) {
		t.Fatal("fail")
	}
}
