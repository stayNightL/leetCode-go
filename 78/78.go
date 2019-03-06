package subsets

import (
	"container/list"
)

func subsets(nums []int) [][]int {
	ans := &[][]int{}
	findSets(nums, list.New(), &ans)
	*ans = append(*ans, []int{})
	return *ans
}
func findSets(nums []int, currut *list.List, ans **[][]int) {
	if len(nums) == 0 {
		return
	}
	for i := 0; i <=len(nums)-1; i++ {
		currut.PushBack(nums[i])
		l := currut.Len()
		a := make([]int, l)
		e := currut.Front()
		j := 0
		for {
			if e != nil {
				a[j] = e.Value.(int)
				e = e.Next()
				j++
			} else {
				break
			}
		}
		**ans = append(**ans, a)

		findSets(nums[i+1:], currut, ans)

		currut.Remove(currut.Back())

	}
}
