package combinationSum

import "sort"

/*
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取。

说明：

所有数字（包括 target）都是正整数。
解集不能包含重复的组合。
示例 1:

输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
示例 2:

输入: candidates = [2,3,5], target = 8,
所求解集为:
[
  [2,2,2,2],
  [2,3,3],
  [3,5]
]
*/
func combinationSum(candidates []int, target int) [][]int {
	ans := &([][]int{})
	sort.Ints(candidates)

	serch(&candidates, target, &ans, []int{}, 0)
	return *ans
}
func serch(candidates *[]int, target int, ans **[][]int, current []int, index int) {
	for i, v := range *candidates {
		if i < index {
			continue
		}
		tmp := target - v
		if tmp > 0 {

			c := make([]int, len(current))
			copy(c, current)
			c = append(c, v)
			serch(candidates, tmp, ans, c, i)

		} else if tmp == 0 {
			c := make([]int, len(current))
			copy(c, current)
			c = append(c, v)
			**ans = append(**ans, c) //这里会返回一个新的地址，但是外层还是旧的
			//current = []int{}
			continue
		} else if tmp < 0 {
			continue
		}
	}
}
