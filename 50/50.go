package myPow

/*
现 pow(x, n) ，即计算 x 的 n 次幂函数。

示例 1:

输入: 2.00000, 10
输出: 1024.00000
示例 2:

输入: 2.10000, 3
输出: 9.26100
示例 3:

输入: 2.00000, -2
输出: 0.25000
解释: 2-2 = 1/22 = 1/4 = 0.25
说明:

-100.0 < x < 100.0
n 是 32 位有符号整数，其数值范围是 [−231, 231 − 1] 。
*/

func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	} else if n == 0 {
		return 1
	}
	if n > 0 {
		return pow(x, x, 1, n)
	} else {
		return 1 / pow(x, x, 1, -n)
	}
}
func pow(y float64, x float64, currut int, n int) float64 {
	temp := n
	if temp >= 2*currut {
		x = x * x
		currut *= 2
		return pow(y, x, currut, n)
	} else if temp-currut == 0 {
		return x
	} else if temp-currut == 1 {
		return x * y
	} else {
		return x * pow(y, y, 1, temp-currut)
	}

}
