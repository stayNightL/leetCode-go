package permute

import "sort"

/*
给定一个可包含重复数字的序列，返回所有不重复的全排列。
*/

func permuteUnique(nums []int) [][]int {
	ans := &[][]int{}
	sort.Ints(nums)
	serch(nums, &ans, []int{})
	return *ans
}
func serch(nums []int, ans **[][]int, curr []int) {

	for i, v := range nums {
		// if i > index {
		// 	continue
		// }
		if len(nums) == 1 {
			curr = append(curr, v)
			**ans = append(**ans, curr)
			return
		}
		if i > 0 && v == nums[i-1] {
			continue
		}
		c := make([]int, len(curr))
		copy(c, curr)
		c = append(c, v)
		n := make([]int, len(nums))
		copy(n, nums)
		if len(nums) > 1 {
			n = append(n[:i], n[i+1:]...)
		}

		// for q, p := range nums {
		// 	if q != i {
		// 		n = append(n, p)
		// 	}
		// }
		serch(n, ans, c)
	}
}
