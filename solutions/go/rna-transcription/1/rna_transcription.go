package strand

import "strings"

func ToRNA(dna string) string {
	dna = strings.ToUpper(strings.TrimSpace(dna))
	complements := map[rune]rune{}
    var sb strings.Builder

    complements['A'] = 'U'
    complements['C'] = 'G'
    complements['G'] = 'C'
    complements['T'] = 'A'

    for _, letter := range dna {
        complement, present := complements[letter]

        if !present {
            return ""
        }
        
        sb.WriteRune(complement)
    }

    return sb.String()
}
