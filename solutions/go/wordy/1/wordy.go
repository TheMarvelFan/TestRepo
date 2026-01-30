package wordy

import (
    "strings"
    "strconv"
    "regexp"
)

func Answer(question string) (int, bool) {
	pattern := regexp.MustCompile(`\bby\b|[^A-Za-z0-9 -]`)
    operands := strings.Split(strings.ToLower(pattern.ReplaceAllString(question, "")), " ")

    if len(operands) < 3 {
        return 0, false
    }

    operands = operands[2:]

    mul := false
    div := false
    add := false
    sub := false
    
    res := 0

    for idx, op := range operands {
        if len(op) == 0 {
            continue
        }

        if idx == 0 {
            num, err := strconv.Atoi(op)

            if err != nil {
                return 0, false
            }

            res = num
            continue
        }

        num, err := strconv.Atoi(op)

        if err != nil {
            switch(op) {
                case "plus":
                	if add || sub || mul || div {
                        return 0, false
                    }

                	add = true
                
                case "minus":
                	if add || sub || mul || div {
                        return 0, false
                    }

                	sub = true

                case "multiplied":
                	if add || sub || mul || div {
                        return 0, false
                    }

                	mul = true

                case "divided":
                	if add || sub || mul || div {
                        return 0, false
                    }

                	div = true

                default:
                	return 0, false
            }
        } else {
            if add {
                res += num
                add = false
            } else if sub {
                res -= num
                sub = false
            } else if mul {
                res *= num
                mul = false
            } else if div {
                res /= num
                div = false
            } else {
                return 0, false
            }
        }
    }

    if add || sub || mul || div {
        return 0, false
    }

    return res, true
}
