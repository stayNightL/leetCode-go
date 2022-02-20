package offer


func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}else {
		h := head
		n := head.Next
		h.Next = deleteNode(n,val)
		return h
	}
}
