package prime

import (
    "math"
    "errors"
)

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n <= 0 {
        return 0, errors.New("there is no zeroth or negative prime number")
    }

    if n == 1 {
		return 2, nil
	}
    
	if n == 2 {
		return 3, nil
	}
	
	limit := estimateNthPrime(n)
	
	return segmentedSieve(n, limit), nil
}

func estimateNthPrime(n int) int {
	fn := float64(n)
	logN := math.Log(fn)
	
	if n < 6 {
		return 15
	}
	
	estimate := fn * (logN + math.Log(logN))
	return int(estimate * 1.3)
}

func segmentedSieve(n, limit int) int {
	sqrtLimit := int(math.Sqrt(float64(limit))) + 1
	smallPrimes := simpleSieve(sqrtLimit)
	
	count := 0
	segmentSize := max(sqrtLimit, 32768)
	
	for low := 0; low <= limit; low += segmentSize {
		high := min(low+segmentSize-1, limit)
		
		segment := make([]bool, high-low+1)
		
		for _, prime := range smallPrimes {
			start := max(prime*prime, (low+prime-1)/prime*prime)
			
			for j := start; j <= high; j += prime {
				segment[j-low] = true
			}
		}
		
		for i := range segment {
			num := low + i
			if num < 2 {
				continue
			}
			
			if num > 2 && num%2 == 0 {
				continue
			}
			
			if !segment[i] {
				count++
				if count == n {
					return num
				}
			}
		}
	}
	
	return -1
}

func simpleSieve(limit int) []int {
	if limit < 2 {
		return []int{}
	}
	
	isPrime := make([]bool, limit+1)
    
	for i := 2; i <= limit; i++ {
		isPrime[i] = true
	}
	
	for i := 2; i*i <= limit; i++ {
		if isPrime[i] {
			for j := i * i; j <= limit; j += i {
				isPrime[j] = false
			}
		}
	}
	
	primes := []int{}
    
	for i := 2; i <= limit; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	
	return primes
}

func max(a, b int) int {
	if a > b {
		return a
	}
    
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
    
	return b
}
