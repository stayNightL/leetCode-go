/*
给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
*/
package maxArea

func maxArea(height []int) int {
	leftIndex := 0
	rightIndex := len(height) - 1
	minSide := 0
	if height[leftIndex] < height[rightIndex] {
		minSide = height[leftIndex]
	} else {
		minSide = height[rightIndex]
	}

	maxArea := minSide * (rightIndex - leftIndex)
	for leftIndex+1 < rightIndex {
		if height[leftIndex] < height[rightIndex] {
			leftIndex++

		} else {
			rightIndex--

		}
		if height[leftIndex] < height[rightIndex] {
			minSide = height[leftIndex]
		} else {
			minSide = height[rightIndex]
		}

		currentArea := minSide * (rightIndex - leftIndex)
		if currentArea > maxArea {
			maxArea = currentArea
		}
	}
	return maxArea
}
