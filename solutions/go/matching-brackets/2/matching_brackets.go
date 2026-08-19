package matchingbrackets

func Bracket(input string) bool {
	strRunes := []rune(input)
	stack := []rune{}

    for _, char := range strRunes {
        switch char {
        case '(', '{', '[':
            stack = append(stack, char)

        case ')':
            if len(stack) > 0 && stack[len(stack) - 1] == '(' {
                stack = stack[ : len(stack) - 1]
            } else {
                return false
            }

        case '}':
            if len(stack) > 0 && stack[len(stack) - 1] == '{' {
                stack = stack[ : len(stack) - 1]
            } else {
                return false
            }

        case ']':
            if len(stack) > 0 && stack[len(stack) - 1] == '[' {
                stack = stack[ : len(stack) - 1]
            } else {
                return false
            }
        }
    }

    return len(stack) == 0
}
