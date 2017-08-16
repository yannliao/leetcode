package leetcode

func search(nums []int, target int) int {
	var result int
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	mid := len(nums) / 2
	last := len(nums) - 1

	if nums[mid] > nums[0] {
		if target == nums[mid] {
			result = mid
		} else if target < nums[mid] && target >= nums[0] {
			result = search(nums[:mid], target) // corner case
		} else {
			if search(nums[mid:], target) != -1 {
				result = mid + search(nums[mid:], target) // corner case
			} else {
				result = -1
			}
		}
	} else if nums[mid] < nums[0] {
		if target == nums[mid] {
			result = mid
		} else if target > nums[mid] && target <= nums[last] {
			if search(nums[mid:], target) != -1 {
				result = mid + search(nums[mid:], target) // corner case
			} else {
				result = -1
			}
		} else {
			result = search(nums[:mid], target) // corner case
		}
	}
	return result
}
