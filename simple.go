package algorithm

import "fmt"

func findIndexesForSum(list []int, target int) []int {
	result := []int{}
	result = append(result, target)

	for index, value := range list {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}

	return result
}
