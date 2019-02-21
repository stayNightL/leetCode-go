/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
*/
package searchRange

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := -1
	right := -1
	leftSerch(&nums, 0, len(nums)-1, target, &left)
	rightSerch(&nums, 0, len(nums)-1, target, &right)

	return []int{left, right}
}

func leftSerch(nums *[]int, min int, max int, target int, left *int) {
	if (*nums)[min] == target {
		*left = min
	} else if (*nums)[max] < target {
		*left = -1
	} else if (*nums)[min] > target {
		*left = -1
	} else if (*nums)[min] < target {
		mid := min + (max-min)/2
		if mid == min {
			if (*nums)[max] == target {
				*left = max
			}
			return
		}
		if (*nums)[mid] > target {
			leftSerch(nums, min, mid, target, left)
		} else if (*nums)[mid] < target {
			leftSerch(nums, mid, max, target, left)
		} else if (*nums)[mid] == target {
			leftSerch(nums, min, mid, target, left)
		}
	}

}
func rightSerch(nums *[]int, min int, max int, target int, right *int) {
	if (*nums)[max] == target {
		*right = max
	} else if (*nums)[max] < target {
		*right = -1
	} else if (*nums)[min] > target {
		*right = -1
	} else if (*nums)[max] > target {
		mid := min + (max-min)/2
		if mid == min {
			if (*nums)[mid] == target {
				*right = mid
			}
			return
		}
		if (*nums)[mid] > target {
			rightSerch(nums, min, mid, target, right)
		} else if (*nums)[mid] < target {
			rightSerch(nums, mid, max, target, right)
		} else if (*nums)[mid] == target {
			rightSerch(nums, mid, max, target, right)
		}
	}

}
