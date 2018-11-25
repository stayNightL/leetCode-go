package getMinimumDifference

import (
	"fmt"
	"testing"
)

func TestGetMini(t *testing.T) {
	var root = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 2,
			},
		},
	}
	fmt.Printf("%d\n", getMinimumDifference(root))
}
