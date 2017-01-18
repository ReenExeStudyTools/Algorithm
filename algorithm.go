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

	accumulator := 0;

	for l.Next != nil && r.Next != nil {
		l = l.Next
		r = r.Next
		sum := accumulator + l.Val + r.Val
		accumulator = 0
		if (sum > 9) {
			accumulator = 1;
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
