/*
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/
package threeSum

import (
	"strconv"
)

type TreeNode struct {
	Val       int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

//FIXME 递归写错了
func (node *TreeNode) toString() string {
	var s string

	if node.LeftNode != nil {
		s += node.LeftNode.toString() + "<" + strconv.Itoa(node.LeftNode.Val)
	}

	if node.RightNode != nil {
		s += strconv.Itoa(node.RightNode.Val) + "<" + node.RightNode.toString()
	}
	return s
}
func (node *TreeNode) add(v int) {
	if node.Val <= v {
		if node.RightNode == nil {
			node.RightNode = &TreeNode{
				Val: v,
			}
		} else {
			node.RightNode.add(v)
		}
	} else {
		if node.LeftNode == nil {
			node.LeftNode = &TreeNode{
				Val: v,
			}
		} else {
			node.LeftNode.add(v)
		}
	}
}

// 构建根节点为0的二叉搜索树
func buildTree(nodes []int) *TreeNode {
	var root = &TreeNode{
		Val: 0,
	}
	for index := range nodes {
		root.add(nodes[index])
	}
	return root
}

// func threeSum(nums []int) [][]int {

// }
