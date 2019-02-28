package getPermutation

import (
	"strconv"
)

/*
给出集合 [1,2,3,…,n]，其所有元素共有 n! 种排列。

按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：

"123"
"132"
"213"
"231"
"312"
"321"
给定 n 和 k，返回第 k 个排列。

说明：

给定 n 的范围是 [1, 9]。
给定 k 的范围是[1,  n!]。
示例 1:

输入: n = 3, k = 3
输出: "213"
示例 2:

输入: n = 4, k = 9
输出: "2314"
*/
func getPermutation(n int, k int) string {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2
	m[3] = 6
	m[4] = 24
	m[5] = 120
	m[6] = 720
	m[7] = 5040
	m[8] = 40320
	m[9] = 362880
	addr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		addr[i] = i
	}
	return getPermutation2(n, k, addr, &m)
}
func getPermutation2(n int, k int, addr []int, m *map[int]int) string {

	if n == 1 {
		return strconv.Itoa(addr[1])
	}
	if k == 0 {
		s := ""
		for _, m := range addr {
			if m != 0 {
				s = strconv.Itoa(m) + s
			}
		}
		return s
	}
	j := k / (*m)[n-1]
	l := k % (*m)[n-1]
	if l != 0 {
		j++
	}
	p := addr[j]
	naddr := append(addr[:j], addr[j+1:]...)
	return strconv.Itoa(p) + getPermutation2(n-1, l, naddr, m)
}
