package offer



func reverseList(head *ListNode) *ListNode {
	//双 指针

	curr := head
	var pre *ListNode =nil
	for curr != nil{
		t := curr.Next
		curr.Next = pre
		pre = curr
		curr = t
	}
	return curr
}
