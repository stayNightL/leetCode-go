package deleteDuplicates

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	l := &ListNode{
		Val:  -1,
		Next: head,
	}

	left := l.Next
	lastNode := l
	if left == nil {
		return left
	}
	right := left.Next
	if right == nil {
		return left
	}
	for i := 0; ; {
		for right != nil {
			if right.Val == left.Val {
				right = right.Next
				i++
				if right == nil {
					break
				}

			} else {
				break
			}
		}
		if i > 0 {
			lastNode.Next = right

		} else {
			lastNode = left

		}
		left = lastNode.Next
		if left == nil {
			break
		}
		right = left.Next
		i = 0
		if right == nil || left == nil {
			break
		}
	}
	return l.Next
}
