package queenattack

import (
    "errors"
    "math"
    "unicode"
)

var ERR_INVALID_POS = errors.New("Invalid Positions")

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if len(whitePosition) != 2 || len(blackPosition) != 2 || whitePosition == blackPosition {
        return false, ERR_INVALID_POS
    }

    charArrWhite := []rune(whitePosition)
    charArrBlack := []rune(blackPosition)

    if !unicode.IsLetter(charArrWhite[0]) || !unicode.IsLetter(charArrBlack[0]) || charArrWhite[0] < 'a' || charArrWhite[0] > 'h' || charArrBlack[0] < 'a' || charArrBlack[0] > 'h' || !unicode.IsDigit(charArrWhite[1]) || !unicode.IsDigit(charArrBlack[1]) || charArrWhite[1] < '1' || charArrWhite[1] > '8' || charArrBlack[1] < '1' || charArrBlack[1] > '8' {
        return false, ERR_INVALID_POS
    }

    return charArrWhite[0] == charArrBlack[0] || charArrWhite[1] == charArrBlack[1] || math.Abs(float64(charArrWhite[1] - charArrBlack[1])) == math.Abs(float64(charArrWhite[0] - charArrBlack[0])), nil
}
