package atbash

import (
    "strings"
    "unicode"
)

func Atbash(s string) string {
	counter := 0
	var sb strings.Builder

    s = strings.ToLower(s)

    for _, char := range s {
        if unicode.IsLetter(char) || unicode.IsDigit(char) {
			if counter == 5 {
                counter = 0
                sb.WriteRune(' ')
            }
            
            if unicode.IsLetter(char) {
                diff := int(char - 'a')
                sb.WriteRune('z' - rune(diff))
            } else {
                sb.WriteRune(char)
            }

            counter ++
        }
    }

    return sb.String()
}
