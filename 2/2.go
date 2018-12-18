/*
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
*/

package addTwoNumbers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	list := &ListNode{
		Val:  0,
		Next: nil,
	}
	head := list
	for {

		if l1 != nil {
			list.Val = l1.Val + list.Val
			l1 = l1.Next

		}

		if l2 != nil {
			list.Val = l2.Val + list.Val
			l2 = l2.Next

		}

		if l1 == nil && l2 == nil {
			break
		}

		list.Next = &ListNode{
			Val:  0,
			Next: nil,
		}
		list = list.Next
	}

	node := head

	for {
		if node.Val >= 10 {
			node.Val = node.Val - 10
			if node.Next != nil {
				node.Next.Val = node.Next.Val + 1
			} else {
				node.Next = &ListNode{
					Val:  1,
					Next: nil,
				}
				break
			}

		}
		fmt.Printf("%d", node.Val)
		node = node.Next
		if node == nil {
			break
		}
	}

	return head
}
