package offer


func getKthFromEnd(head *ListNode, k int) *ListNode {
	h1 :=head
	h2 := h1
	k1 := k
	for  {
		if head.Next == nil {
			break
		}
		k1--
		head = head.Next
		h1 = head
		if k1 <=0 {
			h2 = h2.Next
		}
	}
	return h2

}
