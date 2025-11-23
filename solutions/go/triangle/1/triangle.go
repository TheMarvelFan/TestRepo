// Package triangle provides methods to check the type of triangle.
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
// type Kind
type Kind string

const (
	NaT Kind = "NaT"
	Equ Kind = "Equ"
	Iso Kind = "Iso"
	Sca Kind = "Sca"
)

// KindFromSides returns type of triangle from the sides lengths.
func KindFromSides(a, b, c float64) Kind {
	abEc := a + b
    bcEa := b + c
    caEb := c + a

    if abEc < c || bcEa < a || caEb < b || a <= 0.0 || b <= 0.0 || c <= 0.0 {
        return NaT
    }

    if a == b && b == c {
        return Equ
    }

    if a == b || b == c || c == a {
        return Iso
    }

    return Sca
}
