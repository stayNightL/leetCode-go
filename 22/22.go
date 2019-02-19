package generateParenthesis

import "fmt"

/**
给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
**/
func generateParenthesis(n int) []string {

	r := []string{}
	if n == 0 {
		return append(r, "")
	} else {

		for i := 0; i < n; i++ {
			leftArr := generateParenthesis(i)
			rightArr := generateParenthesis(n - i - 1)
			for _, left := range leftArr {
				for _, right := range rightArr {
					s := "(" + left + ")" + right
					fmt.Println(s)
					r = append(r, s)
				}
			}

		}
	}
	return r
}
