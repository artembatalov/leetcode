func isValid(s string) bool {
	stack := []rune{}
	stack = append(stack, rune(s[0]))
	for i := 1; i < len(s); i++ {
		if s[i] == '{' || s[i] == '[' || s[i] == '(' {
			stack = append(stack, rune(s[i]))
		} else {
			if len(stack) == 0 {
				return false
			} else if stack[len(stack)-1] == '{' && s[i] == '}' {
				stack = stack[:len(stack)-1]
			} else if stack[len(stack)-1] == '[' && s[i] == ']' {
				stack = stack[:len(stack)-1]
			} else if stack[len(stack)-1] == '(' && s[i] == ')' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
