package saddlepoints

import (
    "errors"
    "strconv"
    "strings"
)

// Define the Matrix type here.
type Matrix [][]int

var malformedStrErr = errors.New("Malformed matrix string detected")

func New(s string) (Matrix, error) {
	rowStrs := strings.Split(s, "\n")
	
    if len(rowStrs) == 0 {
        return nil, malformedStrErr
    }

    var row []int

    internalArr := [][]int{}
    width := -1

    for _, rowStr := range rowStrs {
        row = []int{}
		rowSplit := strings.Fields(rowStr)

        if len(rowSplit) == 0 {
            return nil, malformedStrErr
        }
        
        for _, elem := range rowSplit {
            intVal, errIntConv := strconv.Atoi(elem)

            if errIntConv != nil {
                return nil, malformedStrErr
            }
            
            row = append(row, intVal)
        }
        
        if width == -1 {
            width = len(row)
        } else if width != len(row) {
            return nil, malformedStrErr
        }

        internalArr = append(internalArr, row)
    }

    return Matrix(internalArr), nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m Matrix) Cols() [][]int {
    var colArrRow []int
    
	internalArr := [][]int(m)
    colsArr := [][]int{}

    for i := 0; i < len(internalArr[0]); i++ {
        colArrRow = []int{}
        
        for j := 0; j < len(internalArr); j++ {
            colArrRow = append(colArrRow, internalArr[j][i])
        }
        
        colsArr = append(colsArr, colArrRow)
    }

    return colsArr
}

func (m Matrix) Rows() [][]int {
	var rowArrRow []int
    
	internalArr := [][]int(m)
    rowsArr := [][]int{}

    for i := 0; i < len(internalArr); i++ {
        rowArrRow = []int{}
        
        for j := 0; j < len(internalArr[0]); j++ {
            rowArrRow = append(rowArrRow, internalArr[i][j])
        }

        rowsArr = append(rowsArr, rowArrRow)
    }

    return rowsArr
}

func (m Matrix) Set(row, col, val int) bool {
	internalArr := [][]int(m)

    if col >= len(internalArr) || row >= len(internalArr) || col < 0 || row < 0 {
        return false
    }

    internalArr[row][col] = val
    m = Matrix(internalArr)
    return true
}
