package offer


func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Right,root.Left = mirrorTree(root.Left),mirrorTree(root.Right)
	return root
}
