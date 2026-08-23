package robotname

import (
	"crypto/rand"
	"errors"
)

// Robot type
type Robot struct {
	name string
}

var nameSpace = map[string]struct{}{}

const uppercaseAlphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const digits = "1234567890"

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	for {
		if len(nameSpace) == 676000 {
			return "", errors.New("all possible robot names have been used. No more unique names can be generated")
		}

		randLetters, errRandLetters := createCryptoSecureRandAlphaString()
		if errRandLetters != nil {
			return "", errRandLetters
		}

		randNums, errRandNums := createCryptoSecureRandNumString()
		if errRandNums != nil {
			return "", errRandNums
		}

		potName := randLetters + randNums

		if _, exists := nameSpace[potName]; exists {
			continue
		}

		nameSpace[potName] = struct{}{}
		r.name = potName
		return r.name, nil
	}
}

func (r *Robot) Reset() {
	r.name = ""
	if _, err := r.Name(); err != nil {
		panic(err)
	}
}

func createCryptoSecureRandAlphaString() (string, error) {
	result := make([]byte, 2)
	buffer := make([]byte, 2)
	if _, err := rand.Read(buffer); err != nil {
		return "", err
	}

	alphabetLength := len(uppercaseAlphabet)

	for i := range 2 {
		result[i] = uppercaseAlphabet[int(buffer[i])%alphabetLength]
	}

	return string(result), nil
}

func createCryptoSecureRandNumString() (string, error) {
	result := make([]byte, 3)
	buffer := make([]byte, 3)
	if _, err := rand.Read(buffer); err != nil {
		return "", err
	}

	digitLength := len(digits)

	for i := range 3 {
		result[i] = digits[int(buffer[i])%digitLength]
	}

	return string(result), nil
}
