package leetcode

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	var max int
	var minHeight int
	for i < j {
		dis := j - i
		if height[i] > height[j] {
			minHeight = height[j]
			j--
		} else {
			minHeight = height[i]
			i++
		}
		if max < dis*minHeight {
			max = dis * minHeight
		}
	}
	return max
}
