/*

给定一个含有 M x N 个元素的矩阵（M 行，N 列），请以对角线遍历的顺序返回这个矩阵中的所有元素，对角线遍历如下图所示。



示例:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]

输出:  [1,2,4,7,5,3,6,8,9]
*/

package findDiagonalOrder

import (
	"fmt"
)

func findDiagonalOrder(matrix [][]int) []int {
	var currDirect = "up"
	h := len(matrix)
	w := len(matrix[0])
	fmt.Printf("size=====%d", h*w)
	all := w * h
	r := make([]int, all)
	x, y := 0, 0
	for i := 0; all > 0; all-- {
		r[i] = matrix[x][y]
		if currDirect == "up" {
			if isinbond(x-1, y+1, h, w) {
				x--
				y++
				continue
			} else {
				y++
				currDirect = "down"
			}
		} else {
			if isinbond(x+1, y-1, h, w) {
				x++
				y--
				continue
			} else {
				x++
				currDirect = "up"
			}
		}

	}
	return r
}

func isinbond(x int, y int, h int, w int) bool {
	if x < 0 || y < 0 || x > h-1 || y > w-1 {
		return false
	}
	return true
}
