package cryptosquare

import (
    "fmt"
    "math"
    "regexp"
    "strings"
)

func Encode(pt string) string {
	pattern := regexp.MustCompile(`[^a-zA-Z0-9]`)
    formatted := strings.ToLower(pattern.ReplaceAllString(pt, ""))

    length := int(math.Ceil(math.Sqrt(float64(len(formatted)))))
    width := int(math.Floor(math.Sqrt(float64(len(formatted)))))

    builders := make([]strings.Builder, length)

    chars := []rune(formatted)

    for idx, char := range chars {
        builders[idx % length].WriteRune(char)
    }

    padding := width

    if math.Abs(float64(length * length) - float64(len(formatted))) < math.Abs(float64(width * width) - float64(len(formatted))) {
    	padding = length
    }
    
    var ret strings.Builder

    for idx, builder := range builders {
        ret.WriteString(fmt.Sprintf("%-*s", padding, builder.String()))

        if idx < len(builders) - 1 {
            ret.WriteRune(' ')
        }
    }

    return ret.String()
}
