package piglatin

import (
    "slices"
    "strings"
)

var vows = []rune{'a', 'e', 'i', 'o', 'u'}

func Sentence(sentence string) string {
    words := strings.Split(sentence, " ")
    var ret strings.Builder
	var firstTwo string
    var word []rune
    
    for id, wd := range words {
		if strings.TrimSpace(wd) == "" {
            continue
        }
        
        word = []rune(wd)
        firstTwo = wd[0:2]
    	
        if slices.Contains(vows, word[0]) || firstTwo == "xr" || firstTwo == "yt" {
      		ret.WriteString(wd + "ay")
            if id < len(words) - 1 {
            	ret.WriteRune(' ')
        	}
      		continue
    	}

        firstOcc := getFrstVowelOrY(word)
		if firstOcc == -1 {
            ret.WriteString(wd + "ay")
            continue
        }
        
        ret.WriteString(wd[firstOcc:] + wd[0:firstOcc] + "ay")

        if id < len(words) - 1 {
            ret.WriteRune(' ')
        }
	}

    return ret.String()
}

func getFrstVowelOrY(word []rune) int {
    for i := 1; i < len(word); i++ {
        if slices.Contains(vows, word[i]) || word[i] == 'y' {
            if word[i] == 'u' && word[i - 1] == 'q' {
                continue
            } else {
                return i
            }
        }
    }

    return -1
}
