package codewars

import (
	"reflect"
	"testing"
)

func TestFindIndexesForSum(t *testing.T) {
	if !reflect.DeepEqual(RepeatString(1, ""), "") {
		t.Fatal("fail")
	}

	if !reflect.DeepEqual(RepeatString(4, "a"), "aaaa") {
		t.Fatal("fail")
	}
	if !reflect.DeepEqual(RepeatString(3, "hello "), "hello hello hello ") {
		t.Fatal("fail")
	}

	if !reflect.DeepEqual(RepeatString(2, "abc"), "abcabc") {
		t.Fatal("fail")
	}
}

func TestTwoOldestAges(t *testing.T)  {
	if !reflect.DeepEqual(TwoOldestAges([]int{6,5,83,5,3,18}), [2]int{18,83}) {
		t.Fatal("fail")
	}

	if !reflect.DeepEqual(TwoOldestAges([]int{1,5,87,45,8,8}), [2]int{45,87}) {
		t.Fatal("fail")
	}
}

func TestMaxBallTime(t *testing.T)  {
	if MaxBallTime(1) != 1 {
		t.Fatal("fail")
	}
}
