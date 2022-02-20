package offer


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	mirror := mirror(root.Left)
	right := root.Right
	return  equals(mirror,right)

}
func equals(root *TreeNode,B *TreeNode) bool{
	if root == nil && B == nil {
		return true
	}
	if root == nil || B == nil {
		return false
	}
	if root.Val == B.Val {
		return equals(root.Right,B.Right) && equals(root.Left,B.Left)
	}else {
		return false
	}
}
func mirror(root *TreeNode)  *TreeNode {
	if root == nil {
		return root
	}
	root.Left,root.Right = mirror(root.Right),mirror(root.Left)
	return root
}
