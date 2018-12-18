/*

给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
	 请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/
package lengthOfLongestSubstring

import (
	"fmt"
	"strconv"
)

func lengthOfLongestSubstring(s string) int {
	i := 0

	ans := 0
	m := make(map[int32]int)
	for index, value := range s {

		if m[value] == 0 {
			fmt.Println("add" + strconv.QuoteRune(value))
			m[value] = index + 1
			if !(ans > len(m)) {
				ans = len(m)
			}
		} else {

			for _, v := range s[i:m[value]] {
				i++
				delete(m, v)
				fmt.Println("delete" + strconv.QuoteRune(v))
				m[value] = index + 1

			}
			fmt.Println("add" + strconv.QuoteRune(value))

		}

	}
	return ans
}
