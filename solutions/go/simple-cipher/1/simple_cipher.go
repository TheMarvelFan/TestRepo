package cipher

import (
    "strings"
    "unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.

type shift struct {
    key int
}

type vigenere struct {
    key string
}

func NewCaesar() Cipher {
     return vigenere{
        key: "ddd",
    }
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance > 25 || distance == 0 {
        return nil
    }
    
	return shift{
        key: distance,
    }
}

func (c shift) Encode(input string) string {
    input = strings.ToLower(input)
	newLetter := ' '

    var sb strings.Builder
    
    for _, letter := range input {
        if unicode.IsLetter(letter) {
            newLetter = letter + rune(c.key % 25)
    
            for newLetter < 'a' {
                newLetter += rune(26)
            }

            for newLetter > 'z' {
                newLetter -= rune(26)
            }
    
            sb.WriteRune(newLetter)
        }
    }

    return sb.String()
}

func (c shift) Decode(input string) string {
	newLetter := ' '

    var sb strings.Builder
    
    for _, letter := range input {
        if unicode.IsLetter(letter) {
            newLetter = letter - rune((c.key % 25))
    
            for newLetter < 'a' {
                newLetter += rune(26)
            }

            for newLetter > 'z' {
                newLetter -= rune(26)
            }
    
            sb.WriteRune(newLetter)
        }
    }

    return sb.String()
}

func NewVigenere(key string) Cipher {
	countA := 0

    for _, letter := range key {
        if !unicode.IsLetter(letter) || !unicode.IsLower(letter) {
            return nil
        }
        
        if letter == 'a' {
            countA ++
        }
    }

    if countA == len(key) {
        return nil
    }

    strings.TrimSpace(key)
    
	return vigenere{
        key: key,
    }
}

func (v vigenere) Encode(input string) string {
    input = strings.ToLower(input)
	newLetter := ' '

    var sb strings.Builder
    keyId := 0
    
    for _, letter := range input {
        if unicode.IsLetter(letter) {
            newLetter = letter + rune((v.key[keyId] - 'a'))
    
            for newLetter < 'a' {
                newLetter += rune(26)
            }

            for newLetter > 'z' {
                newLetter -= rune(26)
            }
    
            sb.WriteRune(newLetter)

            keyId = (keyId + 1) % len(v.key)
        }
    }

    return sb.String()
}

func (v vigenere) Decode(input string) string {
	newLetter := ' '

    var sb strings.Builder
    keyId := 0
    
    for _, letter := range input {
        if unicode.IsLetter(letter) {
            newLetter = letter - rune((v.key[keyId] - 'a'))
    
            for newLetter < 'a' {
                newLetter += rune(26)
            }

            for newLetter > 'z' {
                newLetter -= rune(26)
            }
    
            sb.WriteRune(newLetter)

            keyId = (keyId + 1) % len(v.key)
        }
    }

    return sb.String()
}
