package coding_tasks

func IsBalancedBrackets(input string) bool {
	if input == "" {
		return true
	}

	var stack ListOfElements

	for _, a := range input {
		switch a {
		case '(':
			stack.Push('(')
		case ')':
			if stack.Peek() == '(' {
				stack.Pop()
			} else {
				return false
			}
		}
	}
	return stack.Tail == nil
}