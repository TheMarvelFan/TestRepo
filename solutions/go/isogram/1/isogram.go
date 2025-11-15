package isogram

import "unicode"

func IsIsogram(word string) bool {
	set := map[rune]struct{}{}

    for _, letter := range word {
        if unicode.IsLetter(letter) {
            letter = unicode.ToLower(letter)
            _, present := set[letter]

            if present {
                return false
            }
            
            set[letter] = struct{}{}
        }
    }

    return true
}
