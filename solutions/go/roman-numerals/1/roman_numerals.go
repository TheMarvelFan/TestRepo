package romannumerals

import (
    "strings"
    "errors"
    "strconv"
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
        return "", errors.New("Input outside of acceptable range.")
    }
    
    var sb strings.Builder
    firstDig := -1
    number := strconv.Itoa(input)
    i := 0
    
	for i < len(number) {
        firstDig = int(rune(number[i]) - '0')
        actual, err := strconv.Atoi(number[i:])

        if err != nil {
            panic(err)
        }

        input = actual

		if input > 1000 {
            for firstDig > 0 {
                sb.WriteString("M")
                firstDig --
            }
        } else if input >= 500 {
            if firstDig == 9 {
                sb.WriteString("CM")
            } else {
                sb.WriteString("D")
                firstDig -= 5

                for firstDig > 0 {
                    firstDig --
                    sb.WriteString("C")
                }
            }
        } else if input >= 100 {
            if firstDig == 4 {
                sb.WriteString("CD")
            } else {
                for firstDig > 0 {
                    firstDig --
                    sb.WriteString("C")
                }
            }
        } else if input >= 50 {
			if firstDig == 9 {
                sb.WriteString("XC")
            } else {
                sb.WriteString("L")
                firstDig -= 5

                for firstDig > 0 {
                    firstDig --
                    sb.WriteString("X")
                }
            }
        } else if input >= 10 {
			if firstDig == 4 {
                sb.WriteString("XL")
            } else {
                for firstDig > 0 {
                    firstDig --
                    sb.WriteString("X")
                }
            }
        } else if input >= 5 {
			if firstDig == 9 {
                sb.WriteString("IX")
            } else {
                sb.WriteString("V")
                firstDig -= 5

                for firstDig > 0 {
                    firstDig --
                    sb.WriteString("I")
                }
            }
        } else {
			if firstDig == 4 {
                sb.WriteString("IV")
            } else {
                for firstDig > 0 {
                    firstDig --
                    sb.WriteString("I")
                }
            }
        }
        
        i ++
    }

    return sb.String(), nil
}
