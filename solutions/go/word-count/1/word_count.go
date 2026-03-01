package wordcount

import (
    "regexp"
    "strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	re := regexp.MustCompile(`\b[\w']+\b`)
    matches := re.FindAllString(phrase, -1)

    wordCount := Frequency{}
	for _, word := range matches {
		lower := strings.ToLower(word)
		wordCount[lower]++
	}

    return wordCount
}
