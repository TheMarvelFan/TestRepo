package luhn

import (
    "unicode"
    "strings"
)

func Valid(id string) bool {
	id = strings.TrimSpace(id)
    
	if len(id) <= 1 {
        return false
    }
    
	sum := 0
    double := false

    for i := len(id) - 1; i >= 0; i -- {
        if unicode.IsDigit(rune(id[i])) {
            add := int(rune(id[i]) - '0')

            if double {
                add *= 2
                
                if add > 9 {
                    add -= 9
                }
            }

            sum += add
            double = !double
        } else {
            if rune(id[i]) != ' ' {
                return false
            }
        }
    }

    return sum % 10 == 0
}
