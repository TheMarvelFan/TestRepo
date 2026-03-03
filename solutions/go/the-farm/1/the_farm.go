package thefarm

import "errors"
import "fmt"

// TODO: define the 'DivideFood' function
func DivideFood(calculator FodderCalculator, numberOfCows int) (float64, error) {
    totalAmount, err := calculator.FodderAmount(numberOfCows)

    if err != nil {
        return 0.0, err
    }

    fatteningFactor, errFatten := calculator.FatteningFactor()

    if errFatten != nil {
        return 0.0, errFatten
    }

    return (totalAmount / float64(numberOfCows)) * fatteningFactor, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(calculator FodderCalculator, numberOfCows int) (float64, error) {
    if numberOfCows <= 0 {
        return 0.0, errors.New("invalid number of cows")
    }

    if calculator == nil {
        return 0.0, errors.New("invalid calculator instance")
    }

    return DivideFood(calculator, numberOfCows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(numberOfCows int) error {
    if numberOfCows < 0 {
    	return &InvalidCowsError{
            numberOfCows: numberOfCows,
            errorMessage: "there are no negative cows",
        }
    }

    if numberOfCows == 0 {
        return &InvalidCowsError{
            numberOfCows: 0,
            errorMessage: "no cows don't need food",
        }
    }

    return nil
}

type InvalidCowsError struct{
    numberOfCows int
    errorMessage string
}

func (err *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", err.numberOfCows, err.errorMessage)
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
