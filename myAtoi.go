package leetcode

func myAtoi(str string) int {
	const intMax = 2147483647
	const intMin = -2147483648

	if len(str) == 0 {
		return 0
	}
	flag := false
	neg := false
	var result int32
outlooper:
	for i := 0; i < len(str); i++ {
		var v byte
		d := str[i]
		switch {
		case d == ' ':
			if flag {
				break outlooper
			}
		case d == '+':
			if flag {
				break outlooper
			}
			flag = true
		case d == '-':
			if flag {
				break outlooper
			} else {
				neg = true
				flag = true
			}
		case '0' <= d && d <= '9':
			v = d - '0'
			if result > intMax/10 || (result == intMax/10 && v > 7) {
				if neg {
					return intMin
				}
				return intMax
			}
			result = result*10 + int32(v)
			flag = true
		default:
			break outlooper
		}
	}
	if neg {
		result *= -1
	}
	return int(result)
}
