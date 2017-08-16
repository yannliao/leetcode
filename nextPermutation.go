package leetcode

import "sort"

func nextPermutation(nums []int) {
	k := len(nums) - 2
	for k >= 0 && nums[k] >= nums[k+1] {
		k--
	}
	if k == -1 {
		sort.Ints(nums)
		return
	}
	var l int
	for i := k + 1; i < len(nums); i++ {
		if nums[i] > nums[k] {
			l = i
		} else {
			break
		}
	}
	nums[l], nums[k] = nums[k], nums[l]
	sort.Ints(nums[k+1:])
}
