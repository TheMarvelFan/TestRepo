package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	ret := map[string]int{}

    for score, letters := range in {
        for _, letter := range letters {
            ret[strings.ToLower(letter)] = score
        }
    }

    return ret
}
