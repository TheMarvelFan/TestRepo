package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	count := 0
    seenBefore := map[int]bool{}
    broken := false
    
	for n != 1 {
		if seenBefore[n] {
            broken = true
            break
        }
        
		seenBefore[n] = true
        
        if n % 2 == 0 {
            n /= 2
        } else {
            n = (n * 3) + 1
        }

        count ++
    }

    if broken {
        return 0, errors.New("Not achievable")
    }

    return count, nil
}
