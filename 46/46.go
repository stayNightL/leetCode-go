package permute

/*
给定一个没有重复数字的序列，返回其所有可能的全排列。
*/

func permute(nums []int) [][]int {
	ans := &[][]int{}

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
