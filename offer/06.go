package offer

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 前序遍历 第一个元素是 根节点
	// 中序遍历 根节点左边是左子树 ，右边是右子树
	if len(preorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	root := preorder[0]
	rt := &TreeNode{Val: root}
	rootindex := 0
	for i, e := range inorder {
		if e == root {
			rootindex = i
			break
		}
	}
	inleft := inorder[:rootindex]
	inRight := inorder[rootindex+1:]
	preLeft := preorder[:len(inleft)+1]
	preRight := preorder[len(inleft)+1:]
	fmt.Println(inleft,inRight,preLeft,preRight)
	rt.Left = buildTree(preLeft,inleft)
	rt.Right = buildTree(preRight,inRight)
	return rt
}

