package leetcode

import "sort"

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}
	sort.Ints(nums)
	result := target - nums[0] - nums[1] - nums[2]
outer:
	for i := 0; i < len(nums)-2; i++ {
		// if i > 0 && nums[i] == nums[i-1] {
		// 	continue
		// }
		j := i + 1
		k := len(nums) - 1
		tar := target - nums[i]
		for j < k {
			if abs(tar-nums[j]-nums[k]) <= abs(result) {
				result = tar - nums[j] - nums[k]
			}
			if nums[j] == tar-nums[k] {
				result = 0
				break outer
			} else if nums[j]+nums[k] < tar {
				j++
			} else {
				k--
			}

		}
	}
	return target - result
}
