package partition

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {

	right := &ListNode{
		Val:  x,
		Next: nil,
	}
	left := &ListNode{
		Val:  x,
		Next: head,
	}
	lhead, rhead := left, right
	for {
		if head != nil {

			if head.Val >= x {
				right.Next = &ListNode{
					Val:  head.Val,
					Next: nil,
				}
				right = right.Next

			} else {
				left.Next = &ListNode{
					Val:  head.Val,
					Next: nil,
				}
				left = left.Next

			}
			head = head.Next
			continue
		}
		break
	}
	left.Next = rhead.Next
	return lhead.Next
}
