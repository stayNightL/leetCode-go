package searchmatrix

/*
编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

每行中的整数从左到右按升序排列。
每行的第一个整数大于前一行的最后一个整数。
示例 1:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 3
输出: true
示例 2:

输入:
matrix = [
  [1,   3,  5,  7],
  [10, 11, 16, 20],
  [23, 30, 34, 50]
]
target = 13
输出: false
*/
func searchMatrix(matrix [][]int, target int) bool {
	colMax := len(matrix[0])
	rowMax := len(matrix)
	if colMax == 0 || rowMax == 0 {
		return false
	}
	var i int
	for i = 0; i < rowMax; i++ {
		if matrix[i][rowMax-1] > target {
			break
		} else if matrix[i][rowMax-1] == target {
			return true
		}
	}
	return findTarget(&matrix[i], target)
}

func findTarget(matrix *[]int, target int) bool {
	begin := 0
	end := len(*matrix) - 1
	if (*matrix)[begin] == target || (*matrix)[end] == target {
		return true
	}
	if end-begin <= 1 {
		return false
	}
	mid := begin + (end-begin)/2
	if (*matrix)[mid] > target {
		matrix := (*matrix)[begin:mid]
		return findTarget(&matrix, target)
	}
	m := (*matrix)[mid:end]
	return findTarget(&m, target)

}
