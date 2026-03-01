package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
    r float64
    i float64
}

func (n Number) Real() float64 {
	return n.r
}

func (n Number) Imaginary() float64 {
	return n.i
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
        n1.Real() + n2.Real(),
        n1.Imaginary() + n2.Imaginary(),
    }
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
        n1.Real() - n2.Real(),
        n1.Imaginary() - n2.Imaginary(),
    }
}

func (n1 Number) Multiply(n2 Number) Number {
	return Number{
        (n1.Real() * n2.Real()) - (n1.Imaginary() * n2.Imaginary()),
        (n1.Imaginary() * n2.Real()) + (n1.Real() * n2.Imaginary()),
    }
}

func (n Number) Times(factor float64) Number {
	return Number{
        n.Real() * factor,
        n.Imaginary() * factor,
    }
}

func (n1 Number) Divide(n2 Number) Number {
	return Number{
        ((n1.Real() * n2.Real()) + (n1.Imaginary() * n2.Imaginary())) / (math.Pow(n2.Real(), 2) + math.Pow(n2.Imaginary(), 2)),
        ((n1.Imaginary() * n2.Real()) - (n1.Real() * n2.Imaginary())) / (math.Pow(n2.Real(), 2) + math.Pow(n2.Imaginary(), 2)),
    }
}

func (n Number) Conjugate() Number {
	return Number{
        n.Real(),
        -n.Imaginary(),
    }
}

func (n Number) Abs() float64 {
	return math.Sqrt(math.Pow(n.Real(), 2) + math.Pow(n.Imaginary(), 2))
}

func (n Number) Exp() Number {
    return Number{
    	math.Pow(math.E, n.Real()) * math.Cos(n.Imaginary()),
        math.Pow(math.E, n.Real()) * math.Sin(n.Imaginary()),
    }
}
