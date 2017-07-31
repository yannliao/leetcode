package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	if len(nums) > 3 {
		for i := 0; i < len(nums)-3; i++ {
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}
			left := target - nums[i]
			temp := threeSumNew(nums[i+1:], left)
			if len(temp) > 0 {
				for _, arr := range temp {
					arr = append([]int{nums[i]}, arr...)
					res = append(res, arr)
				}
			}
		}
	}
	return res
}

func threeSumNew(nums []int, target int) [][]int {
	var result [][]int
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		t := target - nums[i]
		for j < k {
			if nums[j] == t-nums[k] {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				for j < k && nums[k] == nums[k+1] {
					k--
				}
			} else if nums[j] > t-nums[k] {
				k--
			} else {
				j++
			}
		}
	}
	return result
}
