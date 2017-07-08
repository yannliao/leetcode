package leetcode

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	j := 0
outer:
	for {
		for i := 0; i < length; i++ {
			if j >= len(strs[i]) || strs[i][j] != strs[0][j] {
				break outer
			}
		}
		j++
	}
	return strs[0][0:j]
}
