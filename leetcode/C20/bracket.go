package C20

// Show how to use slice to simulate stack in golang.
// push : using append
// pop  : shrink the index
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]rune, 0, len(s))

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			// push into the stack
			stack = append(stack, c)
		} else {
			if c == ')' && (len(stack) == 0 || stack[len(stack)-1 ] != '(') {
				return false
			} else if c == '}' && (len(stack) == 0 || stack[len(stack)-1] != '{') {
				return false
			} else if c == ']' && (len(stack) == 0 || stack[len(stack)-1] != '[') {
				return false
			}

			//pop the stack
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
