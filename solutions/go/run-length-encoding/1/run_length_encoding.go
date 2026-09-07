package runlengthencoding

import (
    "strconv"
    "strings"
    "unicode"
)

func RunLengthEncode(input string) string {
	if removeAllWhitespace(input) == "" {
        return ""
    }
    
    var sb strings.Builder
	strRunes := []rune(input)
	count := 0
    
    for i, char := range strRunes {
        if i > 0 && strRunes[i - 1] != char {
            if count > 1 {
            	sb.WriteString(strconv.Itoa(count))
            }
            
            sb.WriteRune(strRunes[i - 1])
            count = 0
        }

        count ++
    }

    if count > 1 {
        sb.WriteString(strconv.Itoa(count))
    }

    sb.WriteRune(strRunes[len(strRunes) - 1])

    return sb.String()
}

func RunLengthDecode(input string) string {
	if removeAllWhitespace(input) == "" {
        return ""
    }
    
	var sb strings.Builder
	var countStr strings.Builder
	strRunes := []rune(input)

    for _, char := range strRunes {
        if unicode.IsDigit(char) {
            countStr.WriteRune(char)
        } else {
            count, errCount := strconv.Atoi(countStr.String())

            if errCount == nil {
                for count > 0 {
                	count --
                    sb.WriteRune(char)
            	}
            } else {
                sb.WriteRune(char)
            }

            countStr.Reset()
        }
    }

    return sb.String()
}

func removeAllWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	
	for _, r := range s {
		if r != ' ' && r != '\t' && r != '\n' && r != '\r' {
			b.WriteRune(r)
		}
	}
    
	return b.String()
}

