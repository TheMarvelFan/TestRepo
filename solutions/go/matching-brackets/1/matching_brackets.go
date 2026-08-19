package matchingbrackets

func Bracket(input string) bool {
	strRunes := []rune(input)
	stack := []rune{}

    for _, char := range strRunes {
        if char == '(' || char == '{' || char == '[' {
            stack = append(stack, char)
        }

        if char == ')' {
            if len(stack) > 0 && stack[len(stack) - 1] == '(' {
                stack = stack[ : len(stack) - 1]
            } else {
                return false
            }
        }

        if char == '}' {
            if len(stack) > 0 && stack[len(stack) - 1] == '{' {
                stack = stack[ : len(stack) - 1]
            } else {
                return false
            }
        }

        if char == ']' {
            if len(stack) > 0 && stack[len(stack) - 1] == '[' {
                stack = stack[ : len(stack) - 1]
            } else {
                return false
            }
        }
    }

    return len(stack) == 0
}
