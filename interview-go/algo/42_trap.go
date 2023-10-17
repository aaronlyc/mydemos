package algo

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight, ans := 0, 0, 0

	for left <= right {
		if height[left] > maxLeft {
			maxLeft = height[left]
		}
		if height[right] > maxRight {
			maxRight = height[right]
		}

		if maxLeft <= maxRight {
			if height[left] < maxLeft {
				ans += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] < maxRight {
				ans += maxRight - height[right]
			}
			right--
		}
	}

	return ans
}
