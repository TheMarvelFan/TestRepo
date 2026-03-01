package resistorcolortrio

import (
    "fmt"
    "strings"
)

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	refCol := map[string]int{
        "black": 0,
        "brown": 1,
        "red": 2,
        "orange": 3,
        "yellow": 4,
        "green": 5,
        "blue": 6,
        "violet": 7,
        "grey": 8,
        "white": 9,
    }

    var sb strings.Builder

    res := (refCol[colors[0]] * 10) + (refCol[colors[1]])
    extraZero := 0
    
    if refCol[colors[1]] == 0 && refCol[colors[0]] != 0 {
        res /= 10
        extraZero ++
    }

    sb.WriteString(fmt.Sprintf("%v", res))

    if zeros := refCol[colors[2]] + extraZero; zeros == 9 {
        sb.WriteString(" giga")
    } else if zeros < 9 && zeros >= 6 {
        for i := zeros - 6; i > 0; i -- {
            sb.WriteRune('0')
        }
        
        sb.WriteString(" mega")
    } else if zeros < 6 && zeros >= 3 {
        for i := zeros - 3; i > 0; i -- {
            sb.WriteRune('0')
        }
        
        sb.WriteString(" kilo")
    } else {
        for zeros > 0 {
            sb.WriteRune('0')
            zeros --
        }

        sb.WriteRune(' ')
    }

    sb.WriteString("ohms")

    return sb.String()
}
