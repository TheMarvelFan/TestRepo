package flowerfield

import "strings"

// Annotate returns an annotated board
func Annotate(board []string) []string {
	ret := []string{}
    runeArr := [][]rune{}
    count := 0
    
	for _, row := range board {
		runeArr = append(runeArr, []rune(row))
    }

    for i, row := range runeArr {
        var sb strings.Builder

        for j, ca := range row {
			count = 0
            
            if isFlower(ca) {
                sb.WriteRune(ca)
                continue
            }

            count += countDiagonal(i, j, runeArr)
            count += countLinear(i, j, runeArr)

            if count == 0 {
                sb.WriteRune(' ')
            } else {
                sb.WriteRune(rune('0' + count))
            }
        }
        
        ret = append(ret, sb.String())
    }

    return ret
}

func countDiagonal(x int, y int, matrix [][]rune) int {
	count := 0
    
    if x > 0 {
        if y > 0 && isFlower(matrix[x - 1][y - 1]) {
            count ++
        }

        if y < len(matrix[x]) - 1 && isFlower(matrix[x - 1][y + 1]) {
            count ++
        }
    }
    
    if x < len(matrix) - 1 {
        if y > 0 && isFlower(matrix[x + 1][y - 1]) {
            count ++
        }

        if y < len(matrix[x]) - 1 && isFlower(matrix[x + 1][y + 1]) {
            count ++
        }
    }

    return count
}

func countLinear(x int, y int, matrix [][]rune) int {
    count := 0
    
    if x > 0 && isFlower(matrix[x - 1][y]) {
        count ++
    }

    if x < len(matrix) - 1 && isFlower(matrix[x + 1][y]) {
        count ++
    }

    if y < len(matrix[x]) - 1 && isFlower(matrix[x][y + 1]) {
        count ++
    }

    if y > 0 && isFlower(matrix[x][y - 1]) {
        count ++
    }

    return count
}

func isFlower(char rune) bool {
    return char == '*'
}
