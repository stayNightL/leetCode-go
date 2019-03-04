package setzeroes

/*
给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法。

示例 1:

输入:
[
  [1,1,1],
  [1,0,1],
  [1,1,1]
]
输出:
[
  [1,0,1],
  [0,0,0],
  [1,0,1]
]
示例 2:

输入:
[
  [0,1,2,0],
  [3,4,5,2],
  [1,3,1,5]
]
输出:
[
  [0,0,0,0],
  [0,4,5,0],
  [0,3,1,0]
]
进阶:

一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
你能想出一个常数空间的解决方案吗？

*/

func setZeroes(matrix [][]int) {
	rowMax := len(matrix)
	colMax := len(matrix[0])
	row0 := 1
	col0 := 1
	row := matrix[0]
	for i := range row {
		if row[i] == 0 {
			row0 = 0
		}
	}
	for j := rowMax - 1; j >= 0; j-- {

		if matrix[j][0] == 0 {
			col0 = 0
		}
	}
	for i := 0; i < rowMax; i++ {
		for j := 0; j < colMax; j++ {
			//遍历到一个0，将其所在的行列首个元素置零
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0

			}
		}
	}
	//置零
	for i := rowMax - 1; i > 0; i-- {
		for j := colMax - 1; j > 0; j-- {
			//如果其行列首位有一个为0 则为0
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	row = matrix[0]
	if row0 == 0 {
		for i := range row {
			row[i] = 0
		}
	}
	if col0 == 0 {
		for j := rowMax - 1; j >= 0; j-- {

			matrix[j][0] = 0
		}
	}
}
