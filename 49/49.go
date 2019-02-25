package groupanagrams

import (
	"container/list"
)

/*
给定一个字符串数组，将字母异位词组合在一起。字母异位词指字母相同，但排列不同的字符串。

示例:

输入: ["eat", "tea", "tan", "ate", "nat", "bat"],
输出:
[
  ["ate","eat","tea"],
  ["nat","tan"],
  ["bat"]
]
说明：

所有输入均为小写字母。
不考虑答案输出的顺序
*/
type ana struct {
	AnaMap map[rune]int
	Str    string
}

func groupAnagrams(strs []string) [][]string {
	ans := list.New()
	answers := [][]string{}
	for _, v := range strs {
		m := make(map[rune]int)
		for _, r := range v {
			m[r]++
		}
		ans.PushBack(ana{m, v})
	}

	f := ans.Front()
	for {
		if f == nil {
			break
		}
		answer := []string{}
		fv := f.Value.(ana)

		answer = append(answer, fv.Str)
		next := f.Next()
		for {
			if next == nil {
				break
			}
			nv := next.Value.(ana)
			if fv.Compaile1(&nv) {
				answer = append(answer, nv.Str)
				next1 := next
				next = next.Next()
				ans.Remove(next1)
			} else {
				next = next.Next()
			}
		}
		f = f.Next()
		answers = append(answers, answer)
	}
	// front := ans.Front()
	// for front.Next() != nil {
	// 	answer := []string{}
	// 	front = front.Next()
	// 	l := front.Value.(ana)
	// 	next1 := front
	// 	for next1.Next() != nil {
	// 		next := next1.Next()
	// 		next1 = next.Next()
	// 		m := next.Value.(ana)
	// 		if m.compaile1(&l) {
	// 			answer = append(answer, m.Str)
	// 			ans.Remove(next)
	// 		}
	// 	}

	return answers
}
func (a ana) Compaile1(an *ana) bool {
	if len(a.AnaMap) != len(an.AnaMap) {
		return false
	}

	for i, v := range a.AnaMap {
		if v != an.AnaMap[i] {
			return false
		}
	}
	return true

}
