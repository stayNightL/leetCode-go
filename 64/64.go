package minPathSum

/**
给定一个包含非负整数的 m x n 网格，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。

示例:

输入:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
输出: 7
解释: 因为路径 1→3→1→1→1 的总和最小。
**/

func minPathSum(grid [][]int) int {

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			var left, up int = -1, -1
			if j-1 >= 0 {
				left = grid[i][j-1]
			}
			if i-1 >= 0 {
				up = grid[i-1][j]
			}
			if left == -1 && up == -1 {
				continue
			} else if left == -1 && up != -1 {
				grid[i][j] += up
			} else if left != -1 && up == -1 {
				grid[i][j] += left
			} else if left != -1 && up != -1 {
				if left < up {
					grid[i][j] += left
				} else {
					grid[i][j] += up
				}
			}

		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
