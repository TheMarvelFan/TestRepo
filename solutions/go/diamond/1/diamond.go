package diamond

import (
    "errors"
    "strings"
)

func Gen(char byte) (string, error) {
	if char < byte(65) || char > byte(90) {
        return "", errors.New("Invalid character")
    }

    var sb strings.Builder
    
    charInt := int(char - byte(64))
    diff := 0
    
	for i := 1; i < charInt * 2; i ++ {
        if i > charInt {
            diff += 2
        }
        
        for j := 1; j < charInt * 2; j ++ {
            if j == charInt + 1 - (i - diff) || j == charInt + (i - diff) - 1 {
                sb.WriteRune(rune(i - diff + 64))
            } else {
                sb.WriteRune(' ')
            }
        }
        if i < charInt * 2 - 1 {
        	sb.WriteString("\n")
        }
    }

    return sb.String(), nil
}
