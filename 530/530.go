package getMinimumDifference

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	return rootSubCh(root, 9999)
}
func rootSubCh(root *TreeNode, mini int) int {
	var max = root.Val
	if root.Left != nil {
		Node := root.Left
		for {
			if Node != nil && Node.Right != nil {
				max = Node.Right.Val
				Node = Node.Right
			} else if Node != nil && Node.Right == nil {
				max = Node.Val
				break
			} else {
				break
			}
		}
		max = root.Val - max
		if mini > max {
			mini = max
		}

		mini = rootSubCh(root.Left, mini)

	}

	if root.Right != nil {
		Node := root.Right
		for {

			if Node != nil && Node.Left != nil {

				max = Node.Left.Val
				Node = Node.Left
			} else if Node != nil && Node.Left == nil {
				max = Node.Val
				break
			} else {
				break
			}
		}
		max = max - root.Val
		if mini > max {
			mini = max
		}
		mini = rootSubCh(root.Right, mini)

	}

	return mini
}
