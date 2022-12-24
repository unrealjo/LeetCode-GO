package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	current := result
	sum := 0
	for l1 != nil || l2 != nil || sum != 0 {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		// Applying sum mod 10 to get the first digit of sum only
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		// In order to eliminate the first digit
		sum = sum / 10
	}

	return result.Next
}
