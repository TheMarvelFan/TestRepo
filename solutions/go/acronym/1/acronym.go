// Package acronym provides methods to abbreviate long phrases
package acronym

import (
    "strings"
    "regexp"
    "unicode"
)

// Abbreviate takes the first letter of every separated word and returns the acronym
func Abbreviate(s string) string {
    regex := regexp.MustCompile(`[^a-zA-Z']`)
    words := regex.Split(s, -1)
    
    var letters []rune
    var sb strings.Builder

	for _, word := range words {
        if len(word) != 0 {
            letters = []rune(word)
            sb.WriteRune(unicode.ToUpper(letters[0]))
        }
    }
    
	return sb.String()
}
