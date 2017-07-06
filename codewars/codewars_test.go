package codewars

import (
	"reflect"
	"testing"
)

func TestFindIndexesForSum(t *testing.T) {
	if !reflect.DeepEqual(RepeatString(1, ""), "") {
		t.Fail()
	}

	if !reflect.DeepEqual(RepeatString(4, "a"), "aaaa") {
		t.Fail()
	}
	if !reflect.DeepEqual(RepeatString(3, "hello "), "hello hello hello ") {
		t.Fail()
	}

	if !reflect.DeepEqual(RepeatString(2, "abc"), "abcabc") {
		t.Fail()
	}
}

func TestTwoOldestAges(t *testing.T) {
	if !reflect.DeepEqual(TwoOldestAges([]int{6, 5, 83, 5, 3, 18}), [2]int{18, 83}) {
		t.Fail()
	}

	if !reflect.DeepEqual(TwoOldestAges([]int{1, 5, 87, 45, 8, 8}), [2]int{45, 87}) {
		t.Fail()
	}
}

func TestMaxBallTime(t *testing.T) {
	/**
	if MaxBallTime(37) != 10 {
		t.Fail()
	}

	if MaxBallTime(85) != 24 {
		t.Fail()
	}

	if MaxBallTime(136) != 39 {
		t.Fail()
	}
	*/
}

func TestFinancePlanetPlan(t *testing.T) {
	if FinancePlanetPlan(6) != 168 {
		t.Fail()
	}
}

func TestMultiple3And5(t *testing.T) {
	if Multiple3And5(10) != 23 {
		t.Fail()
	}
}

func TestIsTriangle(t *testing.T) {
	if IsTriangle(5, 1, 2) != false {
		t.Fail()
	}

	if IsTriangle(4, 2, 3) != true {
		t.Fail()
	}
}

func TestHasUniqueChar(t *testing.T) {
	if HasUniqueChar("  nAa") != false {
		t.Fail()
	}

	if HasUniqueChar("AaBbC") != true {
		t.Fail()
	}
}

func TestEquableTriangle(t *testing.T) {
	if EquableTriangle(5, 12, 13) != true {
		t.Fail()
	}

	if EquableTriangle(2, 3, 4) != false {
		t.Fail()
	}
}

func TestFindUniq(t *testing.T) {
	if FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}) != float32(2) {
		t.Fail()
	}

	if FindUniq([]float32{0, 0, 0.55, 0, 0}) != float32(0.55) {
		t.Fail()
	}
}

func TestCrossover(t *testing.T) {
	xs, ys := Crossover([]int{}, []int{1, 1, 1, 1, 1}, []int{2, 2, 2, 2, 2})
	if !reflect.DeepEqual(xs, []int{1, 1, 1, 1, 1}) {
		t.Fail()
	}
	if !reflect.DeepEqual(ys, []int{2, 2, 2, 2, 2}) {
		t.Fail()
	}
}

func TestGap(t *testing.T) {
	if !reflect.DeepEqual(Gap(2, 100, 110), []int{101, 103}) {
		t.Fail()
	}
}
