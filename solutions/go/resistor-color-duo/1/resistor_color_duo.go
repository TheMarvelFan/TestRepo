package resistorcolorduo

import "strings"

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	colorCode := map[string]int{
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

    colors[0] = strings.ToLower(colors[0])
    colors[1] = strings.ToLower(colors[1])

    return colorCode[colors[0]] * 10 + colorCode[colors[1]]
}
