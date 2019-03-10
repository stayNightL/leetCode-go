package search

func search(nums []int, target int) bool {
	if len(nums) < 1 {
		return false
	}
	return serch3(nums, target)
}
func serch2(nums []int, target int) bool {
	l := len(nums) - 1
	mid := l / 2
	if l <= 1 {
		return nums[l] == target || nums[0] == target
	}
	if nums[mid] > target {
		return serch2(nums[:mid], target)
	} else if nums[mid] < target {
		return serch2(nums[mid:], target)
	} else {
		return true
	}
}
func serch3(nums []int, target int) bool {
	l := len(nums) - 1
	mid := l / 2
	if l <= 1 {
		return nums[l] == target || nums[0] == target
	}
	if nums[mid] > nums[0] {
		if nums[mid] >= target && nums[0] <= target {
			return serch2(nums[:mid+1], target)
		} else {
			return serch3(nums[mid:], target)
		}
	} else if nums[l] > nums[mid] {
		if nums[mid] <= target && nums[l] >= target {
			return serch2(nums[mid:], target)
		} else {
			return serch3(nums[:mid+1], target)
		}
	} else {
		if nums[l] == target {
			return true
		}
		return serch3(nums[:l], target)

	}
}
