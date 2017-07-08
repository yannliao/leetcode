package leetcode

func convert(s string, numRows int) string {
	zigzag := make([][]rune, numRows)
	var result string
	if numRows == 1 {
		return s
	}
	for i, r := range s {
		reminder := i % (2 * (numRows - 1))
		if reminder > numRows-1 {
			reminder = 2*(numRows-1) - reminder
		}
		zigzag[reminder] = append(zigzag[reminder], r)
	}
	for i := 0; i < numRows; i++ {
		result += string(zigzag[i])
	}
	return result
}
