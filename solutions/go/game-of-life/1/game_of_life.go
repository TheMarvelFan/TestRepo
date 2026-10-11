package gameoflife

const willBeResurrected = 2
const willDie = -1
const isAlive = 1
const isDead = 0
const exactlyThisManyToBeResurrected = 3
var exactlyTheseManyToStayAlive = []int{2, 3}

func Tick(matrix [][]int) [][]int {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			aliveNeighborsCount := -1

			if i > 0 {
				if i < len(matrix)-1 {
					if j > 0 {
						if j < len(matrix[i])-1 {
							aliveNeighborsCount = countAliveNeighbors(matrix, 1, 1, 1, 1, i, j)
						} else {
							aliveNeighborsCount = countAliveNeighbors(matrix, 1, 1, 1, 0, i, j)
						}
					} else {
						aliveNeighborsCount = countAliveNeighbors(matrix, 1, 0, 1, 1, i, j)
					}
				} else {
					if j > 0 {
						if j < len(matrix[i])-1 {
							aliveNeighborsCount = countAliveNeighbors(matrix, 1, 1, 0, 1, i, j)
						} else {
							aliveNeighborsCount = countAliveNeighbors(matrix, 1, 1, 0, 0, i, j)
						}
					} else {
						aliveNeighborsCount = countAliveNeighbors(matrix, 1, 0, 0, 1, i, j)
					}
				}
			} else {
				if j > 0 {
					if j < len(matrix[i])-1 {
						aliveNeighborsCount = countAliveNeighbors(matrix, 0, 1, 1, 1, i, j)
					} else {
						aliveNeighborsCount = countAliveNeighbors(matrix, 0, 1, 1, 0, i, j)
					}
				} else {
					aliveNeighborsCount = countAliveNeighbors(matrix, 0, 0, 1, 1, i, j)
				}
			}

			checkAliveAndMark(matrix, aliveNeighborsCount, i, j)
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == willBeResurrected {
				matrix[i][j] = isAlive
			} else if matrix[i][j] == willDie {
				matrix[i][j] = isDead
			}
		}
	}

	return matrix
}

func checkAliveAndMark(matrix [][]int, aliveNeighborsCount, i, j int) {
	if aliveNeighborsCount == -1 {
		panic("this should never happen")
	}

	if matrix[i][j] == isAlive {
		match := false

		for _, exactlyThisManyToStayAlive := range exactlyTheseManyToStayAlive {
			if exactlyThisManyToStayAlive == aliveNeighborsCount {
				match = true
				break
			}
		}

		if !match {
			matrix[i][j] = willDie
		}
	} else {
		if aliveNeighborsCount == exactlyThisManyToBeResurrected {
			matrix[i][j] = willBeResurrected
		}
	}
}

func countAliveNeighbors(matrix [][]int, minI, minJ, maxI, maxJ, currI, currJ int) int {
	aliveNeighborsCount := 0

	for k := -minI; k <= maxI && currI+k < len(matrix); k++ {
		for l := -minJ; l <= maxJ && currJ+l < len(matrix[currI]); l++ {
			if isSameIndex(k, l) {
				continue
			}

			if matrix[currI+k][currJ+l] == isAlive || matrix[currI+k][currJ+l] == willDie {
				aliveNeighborsCount++
			}
		}
	}

	return aliveNeighborsCount
}

func isSameIndex(i, j int) bool {
	return i == 0 && j == 0
}
