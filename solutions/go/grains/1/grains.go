package grains

import "errors"
import "math"

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
        return 0, errors.New("Please provide a square number within the range 1 to 64, both inclusive.")
    }

    return uint64(math.Pow(2, float64(number - 1))), nil
}

func Total() uint64 {
	return uint64(18446744073709551615)
}
