package kindergartengarden

import (
    "errors"
    "sort"
    "strings"
    "unicode"
)

var ERR_INVALID_DIAG = errors.New("Invalid diagram")
var ERR_DUP_NAME = errors.New("Duplicate Name")

var defaultChildren = []string{
    "Alice",
    "Bob",
    "Charlie",
    "David",
    "Eve",
    "Fred",
    "Ginny",
    "Harriet",
    "Ileana",
    "Joseph",
    "Kincaid",
    "Larry",
}

// Define the Garden type here.
type Garden map[string][]string

// The diagram argument starts each row with a '\n'.  This allows Go's
// raw string literals to present diagrams in source code nicely as two
// rows flush left, for example,
//
//     diagram := `
//     VVCCGG
//     VVCCGG`
//
// If the children argument is empty, use the list of children defined in the instructions.
// If it is not empty, use the given value.

func NewGarden(diagram string, children []string) (*Garden, error) {
	if children == nil {
        children = defaultChildren
    }
    
	garden := map[string][]string{}

    rows := strings.Split(diagram, "\n")

    if len(rows) < 3 {
        return nil, ERR_INVALID_DIAG
    }

    rowOne := []rune(rows[1])
    rowTwo := []rune(rows[2])

    if len(rowOne) != len(rowTwo) || len(rowTwo) % 2 != 0 || len(rowOne) / 2 != len(children) {
        return nil, ERR_INVALID_DIAG
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
            return nil, ERR_DUP_NAME
        }
        
		charOne := rowOne[idx * 2]
        charTwo := rowOne[(idx * 2) + 1]
        charThree := rowTwo[idx * 2]
        charFour := rowTwo[(idx * 2) + 1]

        if unicode.IsLower(charOne) || unicode.IsLower(charTwo) || unicode.IsLower(charThree) || unicode.IsLower(charFour) {
            return nil, ERR_INVALID_DIAG
        }
        
		plantOne, okOne := lookup[charOne]
		plantTwo, okTwo := lookup[charTwo]
		plantThree, okThree := lookup[charThree]
		plantFour, okFour := lookup[charFour]

        if !okOne || !okTwo || !okThree || !okFour {
            return nil, ERR_INVALID_DIAG
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
