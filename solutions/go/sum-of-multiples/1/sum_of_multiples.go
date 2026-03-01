package summultiples

func SumMultiples(limit int, divisors ...int) int {
	set := map[int]struct{}{}
	sum := 0
    val := 0

    for _, divisor := range divisors {
		if divisor == 0 {
            continue
        }
        
        val = divisor
        
        for val < limit {
            set[val] = struct{}{}
            val += divisor
        }
    }

    for i := range set {
        sum += i
    }

    return sum
}
