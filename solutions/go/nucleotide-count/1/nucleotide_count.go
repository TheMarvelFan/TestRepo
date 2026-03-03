package dna

import "errors"
import "unicode"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
// Start by uncommenting the following line:
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
// Start by uncommenting the following line:
type DNA []rune

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
//
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	h := map[rune]int{}

    h['A'] = 0
    h['C'] = 0
    h['G'] = 0
    h['T'] = 0

    for _, letter := range d {
        letterInUpper := unicode.ToUpper(letter)

        if letterInUpper != 'A' && letterInUpper != 'C' && letterInUpper != 'G' && letterInUpper != 'T' {
            return nil, errors.New("invalid sequence")
        }

        count , _ := h[letterInUpper]
		h[letterInUpper] = count + 1
    }
    
	return h, nil
}
