package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(strings.TrimSpace(input))
    record := [26]bool{}

    for _, letter := range input {
        id := letter - 'a'

        if id > -1 && id < 26 {
            record[id] = true
        }
    }

    for _, present := range record {
        if !present {
            return false
        }
    }

    return true;
}
