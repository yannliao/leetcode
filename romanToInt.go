package leetcode

func romanToInt(s string) int {
	trans := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := trans[s[len(s)-1]]
	for i := len(s) - 2; i >= 0; i-- {
		if trans[s[i]] < trans[s[i+1]] {
			sum -= trans[s[i]]
		} else {
			sum += trans[s[i]]
		}
	}
	return sum
}
