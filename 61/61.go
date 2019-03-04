package rotateRight

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	l := 1
	h := head
	tail := head
	if h == nil {
		return nil
	}
	if k == 0 {
		return h
	}
	for {
		if tail.Next == nil {
			break
		}
		tail = tail.Next
		l++
	}

	mod := (k + l) % l

	if mod == 0 {

		return h
	}
	tail.Next = h
	o := l - mod - 1
	for ; o > 0; o-- {
		h = h.Next
	}

	ans := h.Next
	h.Next = nil
	return ans
}
