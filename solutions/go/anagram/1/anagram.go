package anagram

import "strings"

func Detect(subject string, candidates []string) []string {
	target := mapOccurrances(subject)
    ret := []string{}

    for _, candidate := range candidates {
        if compareOccurrances(target, mapOccurrances(candidate)) && strings.ToLower(candidate) != strings.ToLower(subject) {
            ret = append(ret, candidate)
        }
    }

    return ret
}

func mapOccurrances(word string) map[rune]int {
    word = strings.ToLower(word)
	letterMap := map[rune]int{}
    
    for _, letter := range word {
        if _, present := letterMap[letter]; present {
            letterMap[letter] += 1
        } else {
            letterMap[letter] = 0
        }
    }

    return letterMap
}

func compareOccurrances(subject map[rune]int, candidate map[rune]int) bool {
    if len(subject) != len(candidate) {
        return false
    }

    for letter, count := range subject {
        if candCount, present := candidate[letter]; !present || candCount != count {
            return false
        }
    }

    return true
}
