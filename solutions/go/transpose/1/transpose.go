package transpose

import "strings"

func Transpose(input []string) []string {
	if len(input) == 0 {
        return input
    }

    extraSpaces := map[int]int{}
	buffer := []*strings.Builder{}
    maxLen := -1
    
    for i, row := range input {
        rowRunes := []rune(row)
        
        if len(rowRunes) > maxLen {
            clear(extraSpaces)
            maxLen = len(rowRunes)
        }
        
        for j, char := range rowRunes {
            if j == len(buffer) {
                
                buffer = append(buffer, &strings.Builder{})
                
                for k := 0; k < i; k++ {
                    buffer[j].WriteRune(' ')
                }
            }
            
            buffer[j].WriteRune(char)
            extraSpaces[j] = 0
        }

        for k := len(rowRunes); k < maxLen; k++ {
            buffer[k].WriteRune(' ')
            toDel, toDelExists := extraSpaces[k]

            if !toDelExists || toDel == 0 {
                extraSpaces[k] = 1
            } else {
                extraSpaces[k] = toDel + 1
            }
        }
    }

    ret := make([]string, len(buffer))
    maxId := -1
    var strCand string

    for i := 0; i < len(buffer); i++ {
        strCand = buffer[i].String()
        
        if strCand[len(strCand) - 1] != ' ' {
            maxId = i
        } else if i > maxId {
            toDel, _ := extraSpaces[i]
            strCand = string([]rune(strCand)[:len(strCand) - toDel])
        }

        ret[i] = strCand
    }

    return ret
}
