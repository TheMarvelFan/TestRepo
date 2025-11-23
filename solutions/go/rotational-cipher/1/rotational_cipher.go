package rotationalcipher

import (
    "unicode"
    "strings"
)

func RotationalCipher(plain string, shiftKey int) string {
	var sb strings.Builder

    for _, char := range plain {
        sb.WriteString(string(getEncryptedCharacter(char, shiftKey)))
    }

    return sb.String()
}

func getEncryptedCharacter (c rune, k int) rune {
    if !unicode.IsLetter(c) {
        return c;
    }
    
    base := '/'
    
    if unicode.IsLower(c) {
        base = 'a'
    } else {
        base = 'A'
    }
    
    return rune(((int(c) + k - int(base)) % 26 + int(base)));
}
