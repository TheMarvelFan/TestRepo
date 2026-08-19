package house

import "strings"

var stuff = []string{
    "house",
    "malt",
    "rat",
    "cat",
    "dog",
    "cow",
    "maiden",
    "man",
    "priest",
    "rooster",
    "farmer",
    "horse",
}

var stuffSuff = map[int]string{
    0: "that Jack built.",
    5: "with the crumpled horn",
    6: "all forlorn",
    7: "all tattered and torn",
    8: "all shaven and shorn",
    9: "that crowed in the morn",
    10: "sowing his corn",
    11: "and the hound and the horn",
}

var stuffPref = []string{
    "This is the",
    "that lay in the",
    "that ate the",
    "that killed the",
    "that worried the",
    "that tossed the",
    "that milked the",
    "that kissed the",
    "that married the",
    "that woke the",
    "that kept the",
    "that belonged to the",
}

func Verse(v int) string {
	if v < 0 || v > 12 {
        return ""
    }

	var sb strings.Builder
    
	dup := v - 1
    trav := dup

    for dup > -1 {
        if dup == v - 1 {
            sb.WriteString(stuffPref[0])
        } else {
            sb.WriteString(stuffPref[trav])
        }

        sb.WriteString(" ")

        sb.WriteString(stuff[dup])

        suff, suffExists := stuffSuff[dup]

        if suffExists {
            sb.WriteString(" ")
            sb.WriteString(suff)
        }

        if dup > 0 {
        	sb.WriteString("\n")
        }

        if dup < v - 1 {
        	trav--
        }
        
        dup--
    }

    return sb.String()
}

func Song() string {
    var sb strings.Builder
    
	for i := 1; i < 13; i++ {
        sb.WriteString(Verse(i))

        if i < 12 {
        	sb.WriteString("\n\n")
        }
    }

    return sb.String()
}
