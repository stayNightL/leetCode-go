package merge

/*
给出一个区间的集合，请合并所有重叠的区间。

示例 1:

输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:

输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。

*/
/**
 * Definition for an interval.
 * type Interval struct {
 *	   Start int
 *	   End   int
 * }
 */
type Interval struct {
	Start int
	End   int
}

func merge(intervals []Interval) []Interval {
	c := make([]int, 5000)
	ans := []Interval{}
	for _, value := range intervals {
		end := c[value.Start]
		if end <= value.End {
			c[value.Start] = value.End
		}
		if value.Start == 0 && value.End == 0 {
			if c[value.Start] == 0 {
				c[value.Start] = -1
			}
		}
	}

	return merge2(c, ans)
}
func merge2(c []int, ans []Interval) []Interval {
	max := 0
	l := len(c)
	for i := 0; i < l; {
		if c[i] == 0 {
			i++
			continue
		}
		if i == 0 && c[i] == -1 {
			c[i] = 0
		}
		max = c[i]
		j := i
		for ; j <= max; j++ {
			if max < c[j] {
				max = c[j]
			}
		}

		ans = append(ans, Interval{i, max})
		i = j
	}
	return ans
}
