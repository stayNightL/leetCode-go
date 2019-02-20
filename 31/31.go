package nextPermutation

import "fmt"

/*
实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

*/

func nextPermutation(nums []int) {
	length := len(nums)
	step := 0
	for i := length - 1; i > 0; i-- {
		if nums[i-1] < nums[i] {
			if nums[i-1] > nums[length-1] {
				nums[length-1], nums[i-1] = nums[i-1], nums[length-1]
			} else {
				swap(nums, i-1, length-1)
				fmt.Println(nums)
			}

			return
		}
		step++
	}
	var half int = length / 2
	for i := 0; i < half-1; i++ {

		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]

	}
}
func swap(nums []int, begin int, end int) {
	n := begin
	for ; begin < end; begin++ {
		nums[n], nums[begin+1] = nums[begin+1], nums[n]
	}
}
