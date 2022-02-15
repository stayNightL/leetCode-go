package offer

import "container/list"

type ListNode struct {
    Val int
    Next *ListNode
 }

func reversePrint(head *ListNode) []int {
	queue := list.New()
	if head == nil {
		return []int{}
	}
	l := 0
	for true {
		l++
		queue.PushBack(head.Val)
		if head.Next == nil {
			break
		}else {
			head = head.Next
		}
	}
	ints := make([]int,l)
	last := queue.Front()
	for true {
		ints[l-1]=last.Value.(int)
		l =l-1
		if last.Next() == nil {
			break
		}else{
			last = last.Next()
		}
	}
	return ints
}
