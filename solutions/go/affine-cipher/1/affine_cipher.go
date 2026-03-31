package affinecipher

import (
    "errors"
    "math"
    "regexp"
    "strings"
    "unicode"
)

func Encode(text string, a, b int) (string, error) {
	if !checkCoprime(a) {
        return "", errors.New("Please only use keys coprime with 26 (length of english alphabet)")
    }

    re := regexp.MustCompile(`[^a-zA-Z0-9]`)
    text = strings.ToLower(re.ReplaceAllString(text, ""))

    textRunes := []rune(text)
    var sb strings.Builder

    count := 0

    for id, char := range textRunes {
        if unicode.IsLetter(char) {
            sb.WriteRune(rune((((a * int(char - 'a')) + b) % 26) + 'a'))
        } else {
            sb.WriteRune(char)
        }

        count ++

        if count == 5 && id < len(textRunes) - 1 {
            count = 0
            sb.WriteRune(' ')
        }
    }

    return sb.String(), nil
}

func Decode(text string, a, b int) (string, error) {
	if !checkCoprime(a) {
        return "", errors.New("Please only use keys coprime with 26 (length of english alphabet)")
    }

    textRunes := []rune(text)
    var sb strings.Builder

    for _, char := range textRunes {
        if char == ' ' {
            continue
        }

        if unicode.IsLetter(char) {
        	sb.WriteRune(rune((((int(char - 'a') - b) * modInverse(a)) % 26 + 26) % 26 + 'a'))
        } else {
            sb.WriteRune(char)
        }
    }

    return sb.String(), nil
}

func checkCoprime(a int) bool {
    minNum := int(math.Min(float64(a), 26.0))
    
    for i := 2; i <= minNum; i ++ {
        if a % i == 0 && 26 % i == 0 {
            return false
        }
    }

    return true
}

func modInverse(a int) int {
    a %= 26
    
    for i := 1; i < 26; i++ {
        if (a * i) % 26 == 1 {
            return i
        }
    }
    
    return -1
}
