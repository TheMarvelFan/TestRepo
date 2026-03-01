package phonenumber

import (
    "fmt"
    "regexp"
    "errors"
)

var regex = regexp.MustCompile(`\D`) // `[^0-9]`

func Number(phoneNumber string) (string, error) {
    formatted := regex.ReplaceAllString(phoneNumber, "")
    numberRunes := []rune(formatted)

    if len(numberRunes) < 10 || len(numberRunes) > 11 {
        return "", errors.New("Invalid number.")
    } 
    
    if len(numberRunes) == 11 {
        if numberRunes[0] != '1' {
        	return "", errors.New("Invalid number.")
        }
            
        numberRunes = numberRunes[1 : ]
    }

     if numberRunes[0] == '0' || numberRunes[0] == '1' || numberRunes[3] == '0' || numberRunes[3] == '1' {
         return "", errors.New("Invalid number.")
     }

    return string(numberRunes), nil
}

func AreaCode(phoneNumber string) (string, error) {
	formattedNum, err := Number(phoneNumber)

    if err != nil {
        return "", err
    }

    numberRunes := []rune(formattedNum)

    return string(numberRunes[ : 3]), nil
}

func Format(phoneNumber string) (string, error) {
	formattedNum, err := Number(phoneNumber)

    if err != nil {
        return "", err
    }

    numberRunes := []rune(formattedNum)

    return fmt.Sprintf("(%s) %s-%s", string(numberRunes[ : 3]), string(numberRunes[3 : 6]), string(numberRunes[6 : ])), nil
}
