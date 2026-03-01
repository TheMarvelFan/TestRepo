// Package proverb provides a method that generates proverbs based on input.
package proverb

import "fmt"

// Proverb generates a proverb based on input.
func Proverb(rhyme []string) []string {
	ret := []string{}

    if len(rhyme) == 0 {
        return ret
    }

    for i, j := 0, 1; i < j && j < len(rhyme); i, j = i + 1, j + 1 {
        ret = append(ret, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[j]))
    }

    return append(ret, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
}
