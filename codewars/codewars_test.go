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

func TestTwoOldestAges(t *testing.T) {
	if !reflect.DeepEqual(TwoOldestAges([]int{6, 5, 83, 5, 3, 18}), [2]int{18, 83}) {
		t.Fatal("fail")
	}

	if !reflect.DeepEqual(TwoOldestAges([]int{1, 5, 87, 45, 8, 8}), [2]int{45, 87}) {
		t.Fatal("fail")
	}
}

func TestMaxBallTime(t *testing.T) {
	/**
	if MaxBallTime(37) != 10 {
		t.Fatal("fail")
	}

	if MaxBallTime(85) != 24 {
		t.Fatal("fail")
	}

	if MaxBallTime(136) != 39 {
		t.Fatal("fail")
	}
	*/
}

func TestFinancePlanetPlan(t *testing.T) {
	if FinancePlanetPlan(6) != 168 {
		t.Fatal("fail")
	}
}

func TestMultiple3And5(t *testing.T) {
	if Multiple3And5(10) != 23 {
		t.Fatal("fail")
	}
}

func TestIsTriangle(t *testing.T) {
	if IsTriangle(5, 1, 2) != false {
		t.Fatal("fail")
	}

	if IsTriangle(4, 2, 3) != true {
		t.Fatal("fail")
	}
}

func TestHasUniqueChar(t *testing.T) {
	if HasUniqueChar("  nAa") != false {
		t.Fatal("fail")
	}
}
