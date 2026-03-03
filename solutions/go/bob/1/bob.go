// Package bob provides methods to simulate conversation with Bob.
package bob

import "unicode"
import "strings"

// Hey is used to get appropriate responses for a sentence aimed at Bob.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

    if remark == "" {
        return "Fine. Be that way!"
    }

    isQuestion := false
    isYell := true
    hasLetters := false

	for id, char := range remark {
        isQuestion = id == len(remark) - 1 && char == '?'
        if unicode.IsLetter(char) {
            if isYell {
                isYell = unicode.IsUpper(char)
            }
            
            hasLetters = true
        }
    }

    isYell = isYell && hasLetters

    if isQuestion {
        if isYell {
            return "Calm down, I know what I'm doing!"
        } else {
            return "Sure."
        }
    } else {
        if isYell {
            return "Whoa, chill out!"
        } else {
            return "Whatever."
        }
    }
}
