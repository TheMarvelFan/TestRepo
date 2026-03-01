package lsproduct

import (
    "errors"
    "math"
    "unicode"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if span > len(digits) {
        return 0, errors.New("span must be smaller than string length")
    }

    if span < 0 {
        return 0, errors.New("span must not be negative")
    }
    
    prod := int64(1)
    digit := '0'

    for i := 0; i < span; i ++ {
        digit = rune(digits[i])

        if !unicode.IsDigit(digit) {
            return 0, errors.New("digits input must only contain digits")
        }

        prod *= int64(digit - '0')
    }

    largest := int64(math.Max(float64(prod), -1.0))
    prevDigit := int64(-1)

    for i := span; i < len(digits); i ++ {
		digit = rune(digits[i])

        if !unicode.IsDigit(digit) {
            return 0, errors.New("digits input must only contain digits")
        }

        prevDigit = int64(rune(digits[i - span]) - '0')

        if prevDigit == int64(0) {
            prod = int64(1)
            for j := i - 1; j > i - span; j -- {
                prod *= int64(rune(digits[j]) - '0')
            }
        } else {
            prod /= prevDigit
        }

        prod *= int64(digit - '0')

        largest = int64(math.Max(float64(prod), float64(largest)))
    }

    return largest, nil
}
