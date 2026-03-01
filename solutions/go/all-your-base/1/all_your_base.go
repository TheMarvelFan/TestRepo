package allyourbase

import (
    "errors"
    "math"
    "slices"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
        return nil, errors.New("input base must be >= 2")
    }

    if outputBase < 2 {
        return nil, errors.New("output base must be >= 2")
    }

    pos := float64(len(inputDigits)) - 1.0

    res := 0

    for _, dig := range inputDigits {
        if dig < 0 || dig >= inputBase {
            return nil, errors.New("all digits must satisfy 0 <= d < input base")
        }

        res += dig * int(math.Pow(float64(inputBase), pos))
        pos --
    }

    ret := []int{}

    for res > 0 {
        ret = append(ret, res % outputBase)
        res /= outputBase
    }

    slices.Reverse(ret)

    if len(ret) == 0 {
        ret = append(ret, 0)
    }

    return ret, nil
}