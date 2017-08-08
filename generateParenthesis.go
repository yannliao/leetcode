package leetcode

func generateParenthesis(n int) []string {
	var res []string
	place(&res, "", 0, 0, n)
	return res
}
func place(res *[]string, str string, open int, close int, max int) {
	if len(str) == max*2 {
		*res = append(*res, str)
		return
	}
	if open < max {
		place(res, str+"(", open+1, close, max)
	}
	if close < open {
		place(res, str+")", open, close+1, max)
	}
}
