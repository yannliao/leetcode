package leetcode

func letterCombinations(digits string) []string {
	v := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var res []string
	if len(digits) == 0 {
		return res
	}
	res = []string{""}
	for _, value := range digits {
		var temp []string
		num := value - '0'
		if num < 0 || num > 9 {
			break
		}
		chars := v[num]
		for _, vj := range chars {
			for _, vk := range res {
				temp = append(temp, vk+string(vj))
			}
		}
		res = temp
	}
	return res
}
