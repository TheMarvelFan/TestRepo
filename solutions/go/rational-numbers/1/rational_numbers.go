package rationalnumbers

import "math"

type Rational struct {
	numerator, denominator int
}

// Reduce simplifies a Rational, eg changing Rational{4, 2} into Rational{2, 1}.
func (r Rational) Reduce() Rational {
	gcdOfNumDem := gcd(r.numerator, r.denominator)
    
    if gcdOfNumDem != 0 {
        r.numerator /= gcdOfNumDem
        r.denominator /= gcdOfNumDem
    }

    if r.denominator < 0 {
        r.denominator = int(math.Abs(float64(r.denominator)))
        r.numerator = -r.numerator
    }

    return r
}

func (r Rational) Add(s Rational) Rational {
	res := Rational{
        numerator: (r.numerator * s.denominator) + (s.numerator * r.denominator),
        denominator: s.denominator * r.denominator,
    }
    
	return res.Reduce()
}

func (r Rational) Sub(s Rational) Rational {
	res := Rational{
        numerator: (r.numerator * s.denominator) - (s.numerator * r.denominator),
        denominator: s.denominator * r.denominator,
    }
    
	return res.Reduce()
}

func (r Rational) Mul(s Rational) Rational {
	res := Rational{
        numerator: s.numerator * r.numerator,
        denominator: s.denominator * r.denominator,
    }
    
	return res.Reduce()
}

func (r Rational) Div(s Rational) Rational {
	if s.numerator == 0 {
        return Rational{0, 0}
    }

    res := Rational{
        numerator: s.denominator * r.numerator,
        denominator: s.numerator * r.denominator,
    }
    
	return res.Reduce()
}

func (r Rational) Abs() Rational {
	r.numerator = int(math.Abs(float64(r.numerator)))
	r.denominator = int(math.Abs(float64(r.denominator)))

    return r.Reduce()
}

// Compute r ^ power, a rational raised to an int exponent.
func (r Rational) Exprational(power int) Rational {
    if power < 0 {
        num := r.numerator
    	den := r.denominator
    	r.numerator = int(math.Pow(float64(den), math.Abs(float64(power))))
    	r.denominator = int(math.Pow(float64(num), math.Abs(float64(power))))
    } else {
        r.numerator = int(math.Pow(float64(r.numerator), float64(power)))
    	r.denominator = int(math.Pow(float64(r.denominator), float64(power)))
    }
    
    return r.Reduce()
}

// Compute base ^ r, an int raised to a rational.
func (r Rational) Expreal(base int) float64 {
	return math.Pow(math.Pow(float64(base), float64(r.numerator)), 1.0/float64(r.denominator))
}

func gcd(a, b int) int {
    if b == 0 {
        return a
    }
    
    return gcd(b, a % b)
}
