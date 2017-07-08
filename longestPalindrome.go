package leetcode

func longestPalindrome(s string) string {
	var imax, jmax, max int
	r := make([][]int, len(s))
	for i := range r {
		r[i] = make([]int, len(s))
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j <= i; j++ {
			if j == i {
				r[j][i] = 1
			} else {
				if s[j] == s[i] {
					if j == i-1 {
						r[i][j] = 2
					} else if r[i-1][j+1] >= 0 {
						r[i][j] = r[i-1][j+1] + 2
					} else {
						r[i][j] = -1
					}
				} else {
					r[i][j] = -1
				}
			}
			if r[i][j] > max {
				max = r[i][j]
				imax = i
				jmax = j
			}
		}
	}
	// fmt.Println(r)
	// fmt.Println(max)
	return s[jmax : imax+1]
}
