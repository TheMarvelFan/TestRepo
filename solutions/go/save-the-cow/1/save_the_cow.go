package savethecow

import (
	"errors"
	"maps"
	"slices"
	"strings"
)

type Game struct {
	correctWord       map[rune][]int
	currentState      []rune
	remainingAttempts int
	state             string
}

const initialAttemptsForNewGame = 9
const stateLose = "Lose"
const stateWin = "Win"
const stateOngoing = "Ongoing"

func NewGame(word string) *Game {
	if removeAllWhitespace(word) == "" {
		return nil
	}

	wordSoFar, wordMap := buildGame([]rune(word))

	return &Game{
		correctWord:       wordMap,
		currentState:      wordSoFar,
		remainingAttempts: initialAttemptsForNewGame,
		state:             stateOngoing,
	}
}

func (g *Game) Guess(r rune) error {
	if g.state != stateOngoing {
		if g.state == stateLose {
			return errors.New("cannot guess after the game is lost")
		}

		return errors.New("cannot guess after the game is won")
	}

	appearances, _ := g.correctWord[r]

	if len(appearances) == 0 {
		g.remainingAttempts--

		if g.remainingAttempts == -1 {
			g.state = stateLose
		}

		return nil
	}

	for _, appearance := range appearances {
		g.currentState[appearance] = r
	}

	g.correctWord[r] = appearances[:0]

	setGameState(g)
	return nil
}

func (g *Game) MaskedWord() string {
	return string(g.currentState)
}

func (g *Game) RemainingGuesses() int {
	if g.remainingAttempts == -1 {
		return 0
	}

	return g.remainingAttempts
}

func (g *Game) State() string {
	return g.state
}

func buildGame(word []rune) (wordSoFar []rune, wordMap map[rune][]int) {
	wordMap = map[rune][]int{}
	wordSoFar = []rune{}

	for range word {
		wordSoFar = append(wordSoFar, '_')
	}

	for id, letter := range word {
		appearances, _ := wordMap[letter]
		appearances = append(appearances, id)
		wordMap[letter] = appearances
	}

	return
}

func setGameState(g *Game) {
	remaining := false
	eachCharAppearances := slices.Collect(maps.Values(g.correctWord))

	for _, eachCharAppearance := range eachCharAppearances {
		if len(eachCharAppearance) != 0 {
			remaining = true
			break
		}
	}

	if !remaining {
		g.state = stateWin
		return
	}

	if g.remainingAttempts == -1 {
		g.state = stateLose
	} else {
		g.state = stateOngoing
	}
}

func removeAllWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))

	for _, r := range s {
		if r != ' ' && r != '\t' && r != '\n' && r != '\r' {
			b.WriteRune(r)
		}
	}

	return b.String()
}
