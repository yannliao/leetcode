package leetcode

func longestValidParentheses(s string) int {
	var stack []int
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if s[stack[len(stack)-1]] == '(' {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, i)
				}
			}
		} else if s[i] == '(' {
			stack = append(stack, i)
		}
	}
	max := 0
	if len(stack) == 0 {
		max = len(s)
	} else {
		a, b := len(s), 0
		for len(stack) > 0 {
			b = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if a-b-1 > max {
				max = a - b - 1
			}
			a = b
		}
		if a > max {
			max = a
		}
	}

	return max
}
