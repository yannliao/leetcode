package leetcode

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if hasPrefix(haystack[i:], needle) {
			return i
		}
	}
	return -1
}
func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
