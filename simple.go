package algorithm

func findIndexesForSum(list []int, target int) []int {
	for index, value := range list {
		for innerIndex, innerValue := range list {
			if index == innerIndex {
				continue
			}

			if value + innerValue == target {
				return []int{index, innerIndex};
			}
		}
	}

	return []int{}
}
