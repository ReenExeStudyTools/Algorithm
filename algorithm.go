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
func addTwoListNode(l *ListNode, r *ListNode) *ListNode {
	result := &ListNode{l.Val + r.Val, nil}

	next := result;

	for l.Next != nil && r.Next != nil {
		l = l.Next
		r = r.Next
		next.Next = &ListNode{l.Val + r.Val, nil}
		next = next.Next
	}

	return result
}
