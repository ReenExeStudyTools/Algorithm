package codewars

import (
	"strings"
	"sort"
)

func RepeatString(repititions int, value string) string {
	return strings.Repeat(value, repititions)
}

func TwoOldestAges(ages []int) [2]int {
	sort.Ints(ages)

	return [2]int{ages[len(ages) - 2], ages[len(ages) - 1]};
}
