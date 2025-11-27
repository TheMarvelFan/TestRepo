package sieve

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	if limit == 2 {
		return []int{2}
	}
	
	// Only store odd numbers in bit array (indices represent 3, 5, 7, 9, ...)
	// Index i represents number (2*i + 3)
	size := (limit-1)/2 + 1
	bits := make([]uint64, (size+63)/64) // Use uint64 for better performance
	
	primes := make([]int, 0, limit/10) // Pre-allocate approximate capacity
	primes = append(primes, 2)
	
	sqrtLimit := int(float64(limit)*0.5 + 0.5) // Fast sqrt approximation
	for sqrtLimit*sqrtLimit > limit {
		sqrtLimit--
	}
	
	// Sieve odd numbers only
	for i := 0; i < (sqrtLimit-1)/2; i++ {
		if bits[i>>6]&(1<<(i&63)) == 0 { // Check if prime
			p := 2*i + 3
			// Mark multiples starting from p²
			// Convert p² to index: (p*p - 3) / 2
			start := (p*p - 3) / 2
			step := p
			for j := start; j < size; j += step {
				bits[j>>6] |= 1 << (j & 63)
			}
		}
	}
	
	// Collect all primes
	for i := 0; i < size; i++ {
		if bits[i>>6]&(1<<(i&63)) == 0 {
			p := 2*i + 3
			if p <= limit {
				primes = append(primes, p)
			}
		}
	}
	
	return primes
}
