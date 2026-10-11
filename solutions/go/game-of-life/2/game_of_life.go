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
			aliveNeighborsCount := countAliveNeighbors(matrix, i, j)
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

func countAliveNeighbors(matrix [][]int, i, j int) int {
	aliveNeighborsCount := 0

	for k := -1; k < 2 && i+k < len(matrix); k++ {
		for l := -1; l < 2 && i+k > -1 && j+l < len(matrix[i+k]); l++ {
			if isSameIndex(k, l) {
				continue
			}

			if (j+l > -1) && (matrix[i+k][j+l] == isAlive || matrix[i+k][j+l] == willDie) {
				aliveNeighborsCount++
			}
		}
	}

	return aliveNeighborsCount
}

func isSameIndex(i, j int) bool {
	return i == 0 && j == 0
}
