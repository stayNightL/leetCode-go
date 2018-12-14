/**
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/
package removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

//TODO 未测试
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	Fnode := &ListNode{
		Val:  0,
		Next: head,
	}
	frist := Fnode
	second := Fnode
	for i := 0; i <= n; i++ {
		frist = frist.Next
	}
	for {
		if frist == nil {
			second = second.Next.Next
			break
		}
		frist = frist.Next
		second = second.Next
	}
	return Fnode.Next
}
