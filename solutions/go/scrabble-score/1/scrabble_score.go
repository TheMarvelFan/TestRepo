package scrabble

import "unicode"

func Score(word string) int {
	score := 0

    for _, letter := range word {
        score += getRuneScore(letter)
    }

    return score
}

func getRuneScore(letter rune) int {
    val := map[rune]int{}

    val['a'] = 1
    val['b'] = 3
    val['c'] = 3
    val['d'] = 2
    val['e'] = 1
    val['f'] = 4
    val['g'] = 2
    val['h'] = 4
    val['i'] = 1
    val['j'] = 8
    val['k'] = 5
    val['l'] = 1
    val['m'] = 3
    val['n'] = 1
    val['o'] = 1
    val['p'] = 3
    val['q'] = 10
    val['r'] = 1
    val['s'] = 1
    val['t'] = 1
    val['u'] = 1
    val['v'] = 4
    val['w'] = 4
    val['x'] = 8
    val['y'] = 4
    val['z'] = 10

    if unicode.IsUpper(letter) {
        letter = unicode.ToLower(letter)
    }

    ret, _ := val[letter]

    return ret
}
