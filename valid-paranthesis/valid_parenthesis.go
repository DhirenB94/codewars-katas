package main

func main() {
	isValid("()[]{}")
}

func isValid(s string) bool {
	stack := ""

	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = stack + string(v)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			//pop off the last element
			stack = stack[:len(stack)-1]
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
