package findPairs

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	var result = 0
	var setMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		setMap[nums[i]]++

		// if setMap[i] = nil {
		// 	setMap[i]=0
		// } else{
		// 	setMap[i]++
		// }
	}
	if k == 0 {
		for index := range setMap {

			if setMap[index+k] > 1 {
				result++
			}
		}
	} else {
		for index := range setMap {

			if setMap[index+k] > 0 {
				result++
			}
		}
	}
	return result
}
