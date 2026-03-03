package darts

import "math"

func Score(x, y float64) int {
    distance := math.Sqrt(x * x + y * y)
    
    if distance <= 1 {
        return 10
    }
    
    if distance <= 5 {
        return 5
    }
    
    if distance <= 10 {
        return 1
    }
    
    return 0
}
