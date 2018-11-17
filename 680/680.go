/*
给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。
示例 1:
输入: "aba"
输出: True
示例 2:
输入: "abca"
输出: True
解释: 你可以删除c字符。
注意:
字符串只包含从 a-z 的小写字母。字符串的最大长度是50000
*/

package validPalindrome

//TODO 这个实现性能不够
func validPalindrome(s string) bool {
	var flag = true
	flag = isValid(s)
	if flag {

		return true
	}
	for i, _ := range s {

		if isValid(s[0:i] + s[i+1:]) {
			return true
		}

	}
	return false
}
func isValid(s string) bool {
	var length = len(s)
	if length%2 == 0 {
		var leftstr = s[:length/2]
		var rightstr = s[length/2+1:]
		var flag = string(leftstr) == string(rightstr)
		return flag
	} else {
		var leftstr = s[:length/2+1]
		var rightstr = s[length : length/2]
		var flag = string(leftstr) == string(rightstr)
		return flag

	}

}
