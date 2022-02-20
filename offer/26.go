package offer


func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil || A==nil {
		return false
	}

	return SubStructure(A,B) || isSubStructure(A.Left,B) || isSubStructure(A.Right,B)
}


func SubStructure(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil{
		return false
	}
	if A.Val == B.Val {

		return SubStructure(A.Left,B.Left) && SubStructure(A.Right,B.Right)
	}else {
		return false
	}
}