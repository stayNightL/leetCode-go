/*
判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

示例 1:

输入: 121
输出: true
示例 2:

输入: -121
输出: false
解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
示例 3:

输入: 10
输出: false
解释: 从右向左读, 为 01 。因此它不是一个回文数。
*/
package isPalindrome

//执行用时372ms
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	left := x
	right := 0
	for left > right {
		right = right*10 + left%10
		//进行至少一次运算之后right == 0 说明 x 以0 结尾，必然不是回文数
		if right == 0 {
			return false
		}

		left = left / 10

	}
	if left == right || left == right/10 {
		return true
	}
	return false
}
