package codewars

import (
	"math"
	"sort"
	"strconv"
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
	return (a+b > c) && (a+c > b) && (b+c > a)
}

func HasUniqueChar(str string) bool {
	charExistsMap := map[rune]bool{}

	for _, chr := range str {
		if charExistsMap[chr] {
			return false
		}

		charExistsMap[chr] = true
	}

	return true
}

func EquableTriangle(a, b, c int) bool {
	sides := []int{a, b, c}
	sort.Ints(sides)
	a, b, c = sides[0], sides[1], sides[2]
	return a*a+b*b == c*c
}

func FindUniq(arr []float32) float32 {
	uniqueMap := map[float32]int{}

	for _, val := range arr {
		uniqueMap[val] += 1
	}

	for val, count := range uniqueMap {
		if count == 1 {
			return val
		}
	}

	return float32(0)
}

func Crossover(ns []int, xs []int, ys []int) ([]int, []int) {
	return xs, ys
}

func Gap(g, m, n int) []int {
	primes := make([]int, n+1)
	for i := 2; i <= n; i++ {
		if primes[i] == 0 {
			for j := i * i; j <= n; j += i {
				primes[j] = 1
			}
		}
	}

	for i := m; i <= n-g; i++ {
		if primes[i] == 0 && primes[i+g] == 0 {
			return []int{i, i + g}
		}
	}

	return nil
}

func SquareOrSquareRoot(arr []int) []int {
	result := make([]int, len(arr))
	for key, value := range arr {
		sqrt := int(math.Sqrt(float64(value)))
		if sqrt*sqrt == value {
			result[key] = sqrt
		} else {
			result[key] = value * value
		}
	}
	return result
}

func Potatoes(p0, w0, p1 int) int {
	return w0 * (100 - p0) / (100 - p1)
}

func FizzBuzzCuckooClock(time string) string {
	hours, _ := strconv.ParseInt(time[:2], 10, 64)
	minutes, _ := strconv.ParseInt(time[3:], 10, 64)

	if minutes == 0 {
		word := "Cuckoo"
		repeats := int(hours % 12)
		if repeats == 1 {
			return word
		}

		if repeats == 0 {
			repeats = 12
		}

		return word + strings.Repeat(" Cuckoo", repeats-1)
	}

	if minutes == 30 {
		return "Cuckoo"
	}

	if minutes%15 == 0 {
		return "Fizz Buzz"
	}

	if minutes%3 == 0 {
		return "Fizz"
	}

	if minutes%5 == 0 {
		return "Buzz"
	}

	return "tick"
}

func BlackOrWhiteKey(keyPressCount int) string {
	return ""
}
