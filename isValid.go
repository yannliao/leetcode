package leetcode

func isValid(s string) bool {
	var st []rune
	for _, v := range s {
		switch v {
		case '(', '[', '{':
			st = append(st, v)
		case ')', ']', '}':
			if len(st) == 0 {
				return false
			}
			t := st[len(st)-1]
			if v == ')' && t != '(' {
				return false
			}
			if v == ']' && t != '[' {
				return false
			}
			if v == '}' && t != '{' {
				return false
			}
			st = st[:len(st)-1]
		}
	}
	if len(st) > 0 {
		return false
	}
	return true
}
