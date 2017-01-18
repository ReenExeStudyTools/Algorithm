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

// TODO: refactoring
// https://leetcode.com/problems/add-two-numbers/
func addTwoListNode(l *ListNode, r *ListNode) *ListNode {
	accumulator := 0

	sum := accumulator + l.Val + r.Val

	if sum > 9 {
		accumulator = 1
		sum -= 10
	}

	result := &ListNode{sum, nil}

	next := result

	for ((l != nil && l.Next != nil) || (r != nil && r.Next != nil)) {
		if l != nil {
			l = l.Next
		}

		if r != nil {
			r = r.Next
		}

		sum := accumulator

		if l != nil {
			sum += l.Val
		}

		if r != nil {
			sum += r.Val
		}

		accumulator = 0
		if sum > 9 {
			accumulator = 1
			sum -= 10
		}
		next.Next = &ListNode{sum, nil}
		next = next.Next
	}

	if accumulator != 0 {
		next.Next = &ListNode{accumulator, nil}
	}

	return result
}
