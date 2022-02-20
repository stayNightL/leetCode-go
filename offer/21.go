package offer


func exchange(nums []int) []int {
	maxIndex := len(nums)-1
	if maxIndex < 0 {
		return nums
	}

	min := 0

	for {
		if maxIndex-min < 1 {
			break
		}
		if nums[min]&1 == 1 {
			min ++
			continue
		}else {
			if nums[maxIndex]&1 == 1 {
				nums[min],nums[maxIndex] = nums[maxIndex],nums[min]
			}
			maxIndex --
		}
	}
	return nums
}
