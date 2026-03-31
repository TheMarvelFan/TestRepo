package highscores

import "slices"

type HighScores struct{
    ScoresStore []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{
        ScoresStore: scores,
    }
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.ScoresStore
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.ScoresStore[len(s.ScoresStore) - 1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	return slices.Max(s.ScoresStore)
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
	max := -1
    secMax := -1
    thMax := -1

    for _, num := range s.ScoresStore {
        if num > max {
            thMax = secMax
            secMax = max
            max = num
        } else if num > secMax {
            thMax = secMax
            secMax = num
        } else if num > thMax {
            thMax = num
        }
    }

    if secMax == -1 {
        secMax = max
    }

    if thMax == -1 {
        thMax = secMax
    }

    if max == secMax && secMax == thMax {
        return []int{max}
    } else if secMax == thMax {
        return []int{max, secMax}
    }
    
    return []int{max, secMax, thMax}
}
