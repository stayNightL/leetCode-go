package sortcolors

func sortColors(nums []int) {
	l := len(nums)
	if l == 0 || l == 1 {
		return
	}
	p1, p2 := 0, l-1

	for i := 0; p2 >= i; {
		v := nums[i]
		if v == 0 {
			nums[p1], nums[i] = nums[i], nums[p1]

			p1++
			i++
			continue
		}
		if v == 2 {
			nums[p2], nums[i] = nums[i], nums[p2]
			p2--
			continue
		}

		i++
	}

}
