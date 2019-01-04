/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
*/
package longestCommonPrefix

type runeNode struct {
	R     rune
	Count int
	Next  *runeNode
}
//TODO 这种解法需要考虑的特殊情况太多
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	head := &runeNode{
		R:     0,
		Count: 0,
		Next:  nil,
	}
	for _, value := range strs {
		currentNode := head
		lastNode := head
		for _, r := range value {
			if currentNode == nil {
				currentNode = &runeNode{
					R:     r,
					Count: 1,
					Next:  nil,
				}
				lastNode.Next = currentNode
				continue
			} else if currentNode.R == r {
				currentNode.Count++
				lastNode = currentNode
				currentNode = currentNode.Next
			} else if currentNode.Count == 0 {
				currentNode.R = r
				currentNode.Count++
				lastNode = currentNode
				currentNode = currentNode.Next
			} else if currentNode.R != r {
				currentNode.Count++
				lastNode = currentNode
				currentNode = currentNode.Next
			}

		}
	}
	s := ""

	for head != nil {

		if head.Count > 1 {
			s += string(head.R)
		}

		head = head.Next

	}
	return s
}
