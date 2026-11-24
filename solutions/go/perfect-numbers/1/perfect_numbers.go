package perfectnumbers

import "errors"

// Define the Classification type here.
type Classification string

const (
    ClassificationDeficient = Classification("deficient")
    ClassificationPerfect = Classification("perfect")
    ClassificationAbundant = Classification("abundant")
    ClassificationErr = Classification("")
)

var ErrOnlyPositive = errors.New("number should be greater than 0")

func Classify(n int64) (Classification, error) {
	if n <= int64(0) {
        return ClassificationErr, ErrOnlyPositive
    }

	divsSum := getDivisorsSum(n)

    if divsSum > n {
        return ClassificationAbundant, nil
    }

    if divsSum < n {
        return ClassificationDeficient, nil
    }

    return ClassificationPerfect, nil
}

func getDivisorsSum(n int64) int64 {
    sum := int64(0)
    num := int64(1)

    for num * num <= n {
        if n % num == 0 {
            if num != (n / num) && (n / num) != n {
                sum += n / num
            }

            if num < n {
             	sum += num
            }
        }

        num++
    }
    
    return sum
}
