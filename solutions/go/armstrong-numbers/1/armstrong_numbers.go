package armstrong

import "math"

func IsNumber(n int) bool {
	if n == 0 {
        return true
    }
    
    tmp := n
	sum := 0.0
    count := 0.0
    digits := []int{}

    for tmp > 0 {
        count ++
        digits = append(digits, tmp % 10)
        tmp /= 10
    }

    for _, digit := range digits {
        sum += math.Pow(float64(digit), count)
    }

    return int(sum) == n
}
