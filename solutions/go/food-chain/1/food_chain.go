package foodchain

import (
    "strings"
    "fmt"
)

var ANIMAL_PROPS = map[int]string{
    2: "How absurd to swallow a bird!\n",
    3: "Imagine that, to swallow a cat!\n",
    4: "What a hog, to swallow a dog!\n",
    5: "Just opened her throat and swallowed a goat!\n",
}
var ANIMALS = []string{"fly", "spider", "bird", "cat", "dog", "goat", "cow", "horse"}

const SPIDER_PROP = "wriggled and jiggled and tickled inside her"
const SWALLOW_WHAT = "She swallowed the"
const SWALLOW = "she swallowed"
const SWALLOW_FOR = "to catch the"
const PERHAPS_LINE = "Perhaps she'll die."
const DEAD_LINE = "She's dead, of course!"
const OLD_LADY = "I know an old lady who swallowed a"
const KNOW_NOT = "I don't know"

func Verse(v int) string {
	if v > 8 || v < 1 {
        return ""
    }
    
    v -= 1
    
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%s %s.\n", OLD_LADY, ANIMALS[v]))

    if v == 6 {
        sb.WriteString(fmt.Sprintf("%s how %s a %s!\n", KNOW_NOT, SWALLOW, ANIMALS[v]))
    } else if v == 1 {
        sb.WriteString(fmt.Sprintf("It %s.\n", SPIDER_PROP))
    } else if v > 0 {
        prop, exists := ANIMAL_PROPS[v]

        if !exists {
            sb.WriteString(DEAD_LINE)
            return sb.String()
        }
        
        sb.WriteString(prop)
    }
    
    for v > 0 {
		sb.WriteString(fmt.Sprintf("%s %s %s %s", SWALLOW_WHAT, ANIMALS[v], SWALLOW_FOR, ANIMALS[v - 1]))
        
        if v == 2 {
            sb.WriteString(fmt.Sprintf(" that %s", SPIDER_PROP))
        }
        
		sb.WriteString(".\n")
        v -= 1
    }

    sb.WriteString(fmt.Sprintf("%s why %s the fly. %s", KNOW_NOT, SWALLOW, PERHAPS_LINE))
    return sb.String()
}

func Verses(start, end int) string {
	var sb strings.Builder
    
    for start <= end {
        sb.WriteString(Verse(start))
        
        if start < end {
            sb.WriteString("\n\n")
        }
        
        start ++
    }

    return sb.String()
}

func Song() string {
	return Verses(1, 8)
}
