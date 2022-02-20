package offer

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = l1
	var tail *ListNode = l2
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		head = l2
		tail = l1
	}
	t := head
	for  {
		if head.Next != nil {
			if tail.Val <= head.Next.Val {
				next := head.Next
				head.Next = tail
				tail_next := tail.Next
				tail.Next = next
				tail = tail_next

			}else {
				head = head.Next
			}
		}else {
			head.Next =tail
		}
		if tail == nil {
			break
		}
	}
	return t
}

