package bottlesong

import (
    "fmt"
    "strings"
)

func Recite(startBottles, takeDown int) []string {
	ret := []string{}
    limit := startBottles - takeDown
    
	wordMap := map[int]string{
        0: "no",
        1: "One",
        2: "Two",
        3: "Three",
        4: "Four",
        5: "Five",
        6: "Six",
        7: "Seven",
        8: "Eight",
        9: "Nine",
        10: "Ten",
    }

    bottleOrBottles := map[bool]string {
        true: "bottles",
        false: "bottle",
    }
    
	for startBottles > limit {
        if startBottles - takeDown < limit {
			ret = append(ret, "")
        }
        
        for i := 1; i < 3; i ++ {
			ret = append(ret, fmt.Sprintf("%s green %s hanging on the wall,", wordMap[startBottles], bottleOrBottles[bottleOrBottlesFunc(startBottles)]))
        }

        ret = append(ret, fmt.Sprintf("And if one green bottle should accidentally fall,"))

        startBottles -= 1

        ret = append(ret, fmt.Sprintf("There'll be %s green %s hanging on the wall.", strings.ToLower(wordMap[startBottles]), bottleOrBottles[bottleOrBottlesFunc(startBottles)]))
    }

    return ret
}

func bottleOrBottlesFunc(num int) bool {
    return num != 1 
}
