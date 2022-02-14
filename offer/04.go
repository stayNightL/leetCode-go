package offer

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	//从右上开始遍历
	//如果相等，返回true
	// 如果 target大，去一行，否则去一列
	//重复
	for{
		y := len(matrix)
		if y == 0 {
			return false
		}
		x := len(matrix[0])
		if x == 0  {
			return false
		}
		i := matrix[0][x-1]
		if i == target {
			return true
		} else if i<target {
			matrix = matrix[1:]
		}else if i>target {
			for j := 0; j < y; j++ {
				matrix[j]=matrix[j][:x-1]
			}
		}
		fmt.Println(matrix)
	}


}