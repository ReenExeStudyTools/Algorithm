package algorithm

// https://leetcode.com/problems/two-sum/
func findIndexesForSum(list []int, target int) []int {
	for index, value := range list {
		for innerIndex, innerValue := range list {
			if index == innerIndex {
				continue
			}

			if value+innerValue == target {
				return []int{index, innerIndex}
			}
		}
	}

	return []int{}
}

// https://leetcode.com/problems/add-two-numbers/
func addTwoListNode(l1 *ListNode, l2 *ListNode) *ListNode {
	return &ListNode{1, nil}
}
