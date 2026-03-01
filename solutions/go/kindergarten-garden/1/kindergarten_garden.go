package kindergarten

import (
    "errors"
    "sort"
    "strings"
    "unicode"
)

// Define the Garden type here.
type Garden map[string][]string

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`

func NewGarden(diagram string, children []string) (*Garden, error) {
	garden := map[string][]string{}

    rows := strings.Split(diagram, "\n")

    if len(rows) < 3 {
        return nil, errors.New("Invalid diagram")
    }

    rowOne := []rune(rows[1])
    rowTwo := []rune(rows[2])

    if len(rowOne) != len(rowTwo) || len(rowTwo) % 2 != 0 || len(rowOne) / 2 != len(children) {
        return nil, errors.New("Invalid diagram")
    }

    lookup := map[rune]string{
        'G': "grass",
        'C': "clover",
        'R': "radishes",
        'V': "violets",
    }

    dupChildren := make([]string, len(children))
    copy(dupChildren, children)
    
    sort.Strings(dupChildren)

    for idx, child := range dupChildren {
		_, already := garden[child]

        if already {
            return nil, errors.New("Duplicate Name")
        }
        
		charOne := rowOne[idx * 2]
        charTwo := rowOne[(idx * 2) + 1]
        charThree := rowTwo[idx * 2]
        charFour := rowTwo[(idx * 2) + 1]

        if unicode.IsLower(charOne) || unicode.IsLower(charTwo) || unicode.IsLower(charThree) || unicode.IsLower(charFour) {
            return nil, errors.New("Invalid diagram")
        }
        
		plantOne, okOne := lookup[charOne]
		plantTwo, okTwo := lookup[charTwo]
		plantThree, okThree := lookup[charThree]
		plantFour, okFour := lookup[charFour]

        if !okOne || !okTwo || !okThree || !okFour {
            return nil, errors.New("Invalid diagram")
        }
        
        garden[child] = []string{
            plantOne,
            plantTwo,
            plantThree,
            plantFour,
        }
    }

    createdGarden := Garden(garden)

    return &createdGarden, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := (*g)[child]

    return plants, ok
}
