package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	return []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	code := map[string]int{}
    code["black"] = 0
    code["brown"] = 1
    code["red"] = 2
    code["orange"] = 3
    code["yellow"] = 4
    code["green"] = 5
    code["blue"] = 6
    code["violet"] = 7
    code["grey"] = 8
    code["white"] = 9

    val, ok := code[color]

    if !ok {
        return -1
    }

    return val
}
