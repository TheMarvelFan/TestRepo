package microblog

// import "unicode/utf16"

func Truncate(phrase string) string {
	if (len(phrase) <= 5) {
        return phrase
    }
    
	// specialChars := utf16.Encode([]rune(phrase))
 //    runes := []rune(string(utf16.Decode(specialChars)))
    return string([]rune(phrase)[: 5])
}
