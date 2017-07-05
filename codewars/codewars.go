package codewars

import (
	"sort"
	"strings"
)

func RepeatString(repititions int, value string) string {
	return strings.Repeat(value, repititions)
}

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)

	return [2]int{ages[len(ages)-2], ages[len(ages)-1]}
}

func MaxBallTime(v int) int {
	// h = v*t - 0.5*g*t*t
	// differential v - gt
	// max h when v = gt
	return v
}

func FinancePlanetPlan(n int) int {
	sum := 0
	from := 0
	to := 6
	distance := 7
	for week := 0; week <= n; week++ {
		if week < 6 {
			sum += (from + to) * distance / 2
			distance--
			from += 2
			to++
		} else {
			sum += to
			to++
		}
	}
	return sum
}

func Multiple3And5(number int) int {
	sum := 0

	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}

func IsTriangle(a, b, c int) bool {
	return false
}
