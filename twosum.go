package leetcode

func twoSum(nums []int, target int) []int {
	var result []int
    temp := make(map[int]int)
	for i, v := range nums {
		if ndx, ok := temp[target - v]; !ok {
			temp[v] = i 
		} else {
			result = append(result, ndx, i)
			break;
		}
	}
	return result
}