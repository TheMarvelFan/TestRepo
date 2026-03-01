package lineup

import "fmt"

func Format(name string, number int) string {
    suffix := "th"
	lastDig := number % 10
    lastNum := number % 100

    if lastDig == 1 && lastNum != 11 {
        suffix = "st"
    }

    if lastDig == 2 && lastNum != 12 {
        suffix = "nd"
    }

    if lastDig == 3 && lastNum != 13 {
        suffix = "rd"
    }

    return fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!", name, number, suffix)
}
