package leetcode

func isValid(s string) bool {
	var stack []string

	for _, c := range s {
		if leftChar(string(c)) {
			stack = append(stack, string(c))
		} else {
			if len(stack) == 0 {
				return false
			} else {
				last := stack[len(stack)-1]
				pair := isPair(last, string(c))
				if pair {
					stack = stack[0 : len(stack)-1]
				} else {
					return false
				}
			}
		}
	}

	return len(stack) == 0
}

func leftChar(c string) bool {
	return c == "(" || c == "[" || c == "{"
}

func isPair(left, right string) bool {
	switch left {
	case "(":
		return right == ")"
	case "[":
		return right == "]"
	case "{":
		return right == "}"
	default:
		return false
	}
}
