package eliudseggs

func EggCount(displayValue int) int {
	bitCount := 0
    
	for displayValue > 0 {
        if displayValue % 2 == 1 {
            bitCount ++
        }

        displayValue /= 2
    }

    return bitCount
}
