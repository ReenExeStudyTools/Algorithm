package codewars

import (
	"reflect"
	"testing"
)

func TestFindIndexesForSum(t *testing.T) {
	if !reflect.DeepEqual(RepeatString(1, ""), "") {
		t.Fatal("fail")
	}
}
