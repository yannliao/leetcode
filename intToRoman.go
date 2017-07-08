package leetcode

import "strings"

func intToRoman(num int) string {
	var result []string
	trans := [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "M", "MM", "MMM"},
	}
	result = append(result, trans[3][(num/1000)%10])
	result = append(result, trans[2][(num/100)%10])
	result = append(result, trans[1][(num/10)%10])
	result = append(result, trans[0][(num/1)%10])
	return strings.Join(result, "")
}
