package saddlepoints

import (
    "math"
    "strconv"
    "strings"
)

// Define the Matrix and Pair types here.
type Matrix [][]int
type Pair [2]int

func New(s string) (*Matrix, error) {
	if removeAllWhitespace(s) == "" {
        ret := Matrix([][]int{})
        return &ret, nil
    }

    mat := [][]int{}

    for _, rowStr := range strings.Split(s, "\n") {
        row := []int{}
        
        for _, val := range strings.Split(rowStr, " ") {
            num, convErr := strconv.Atoi(val)

            if convErr != nil {
                return nil, convErr
            }
            
            row = append(row, num)
        }

        mat = append(mat, row)
    }

    ret := Matrix(mat)

    return &ret, nil
}

func (m *Matrix) Saddle() []Pair {
	if m == nil {
        return nil
    }
    
	mat := [][]int(*m)

    if len(mat) == 0 {
        return nil
    }
    
    ret := []Pair{}
	maxPerRow := map[int]int{}
    var min int
    var minRows []int

    for j := 0; j < len(mat[0]); j ++ {
        min = math.MaxInt
        minRows = []int{}
        
        for i := 0; i < len(mat); i++ {
            if mat[i][j] < min {
                min = mat[i][j]
                minRows = minRows[:0]
                minRows = append(minRows, i)
            } else if mat[i][j] == min {
                minRows = append(minRows, i)
            }
        }

        for _, minRow := range minRows {
            maxTree, exists := maxPerRow[minRow]

            if !exists {
                maxOfRow := getRowMax(mat, minRow)

                if maxOfRow == min {
                    ret = append(ret, Pair([2]int{ minRow + 1, j + 1 }))
                }

                maxPerRow[minRow] = maxOfRow
            } else {
                if maxTree == min {
                    ret = append(ret, Pair([2]int{ minRow + 1, j + 1 }))
                }
            }
        }
    }

    return ret
}

func getRowMax(mat [][]int, row int) int {
    if len(mat) <= row {
        panic("Invalid call")
    }

    max := math.MinInt

    for _, val := range mat[row] {
        if max < val {
            max = val
        }
    }

    return max
}

func removeAllWhitespace(s string) string {
	var b strings.Builder
	b.Grow(len(s))
	
	for _, r := range s {
		if r != ' ' && r != '\t' && r != '\n' && r != '\r' {
			b.WriteRune(r)
		}
	}
    
	return b.String()
}
