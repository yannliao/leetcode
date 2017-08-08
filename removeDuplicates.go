package leetcode

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	id := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[id] = nums[i]
			id++
		}
	}
	return id
}
