package series

import "strings"

func All(n int, s string) []string {
	if n <= 0 || n > len(s) {
		return nil
	}

	if n == len(s) {
		return []string{s}
	}

	var sb strings.Builder
	ret := []string{}
	characters := []rune(s)

	for _, character := range characters {
		if len(sb.String()) >= n {
			ret = append(ret, sb.String())
			sb.Reset()
			sb.WriteString(ret[len(ret)-1][1:])
		}

		sb.WriteRune(character)
	}

	return append(ret, sb.String())
}

func UnsafeFirst(n int, s string) string {
	if n <= 0 || n > len(s) {
		return ""
	}

	return s[:n]
}

func First(n int, s string) (string, bool) {
	if n <= 0 || n > len(s) {
		return "", false
	}

	return s[:n], true
}