package containerwithmostwater

func maxArea(height []int) int {
	max, area, left, right, Width := 0, 0, 0, len(height)-1, 0
	for left < right {
		Width = right - left
		if height[left] >= height[right] {
			area = height[right] * Width
			right--
		} else {
			area = height[left] * Width
			left++
		}
		if area > max {
			max = area
		}
	}
	return max
}
