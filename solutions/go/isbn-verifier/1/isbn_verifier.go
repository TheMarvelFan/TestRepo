package isbn

import "unicode"

func IsValidISBN(isbn string) bool {
	lim := 10
    sum := 0

    for id, code := range isbn {
        if lim <= 0 {
            return false
        }
        
        if unicode.IsDigit(code) {
            sum += int((code - '0')) * lim
            lim --
        } else if unicode.IsLetter(code) {
            if id < len(isbn) - 1 {
                return false
            }
            
            if code == 'X' {
                sum += 10 * lim
                lim --
			} else {
                return false
            }
        }
    }

    if lim > 0 {
        return false
    }

    return sum % 11 == 0
}
