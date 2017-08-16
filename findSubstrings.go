package leetcode

func findSubstring(s string, words []string) []int {
	var res []int

	// temp := make(map[string]bool)
	long := len(words)
	length := len(words[0])
	temp := make(map[string]int)
	for _, value := range words {
		if _, ok := temp[value]; ok {
			temp[value]++
		} else {
			temp[value] = 1
		}
	}
	if long == 0 && length == 0 {
		res = append(res, 0)
		return res
	}
	if len(s) < long*length {
		return res
	}
	for i := 0; i <= len(s)-long*length; i++ {
		copy := make(map[string]int)
		for key := range temp {
			copy[key] = temp[key]
		}

		for j := 0; j < long; j++ {
			str := s[i+j*length : i+j*length+length]
			if _, ok := copy[str]; ok {
				if copy[str] == 1 {
					delete(copy, str)
				} else {
					copy[str]--
				}
				if len(copy) == 0 {
					res = append(res, i)
					break
				}
			}
		}
	}
	return res
}
