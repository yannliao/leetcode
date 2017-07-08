package leetcode

func lengthOfLongestSubstring(s string) int {
	myMap := make(map[rune]int)
	j := 0
	count := 0
	for i, v := range s {
		ndx, ok := myMap[v]
		if ok {
			if j > ndx+1 {
				j = ndx + 1
			}
		}
		myMap[v] = i
		if i-j+1 > count {
			count = i-j+1
		}
	}
	return count
}
