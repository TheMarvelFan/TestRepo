package logs

import "fmt"
import "unicode/utf8"
import "strings"

// Application identifies the application emitting the given log.
func Application(log string) string {
	runeMap := make(map[string]string)

    runeMap["U+2757"] = "recommendation"
    runeMap["U+1F50D"] = "search"
    runeMap["U+2600"] = "weather"
    
	for _, char := range log {
        uni := fmt.Sprintf("%U", char)
        val, ok := runeMap[uni]

        if ok {
            return val
        }
    }

    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, string(oldRune), string(newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
