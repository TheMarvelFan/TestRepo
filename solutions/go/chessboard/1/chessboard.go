package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
    
	val, ok := cb[file]

    if !ok {
        return 0
    }

    for _, occupied := range val {
        if occupied {
            count += 1
        }
    }

    return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
        return 0
    }
    
	count := 0
    rank -= 1
    
	for _, file := range cb {
		if len(file) > rank {
            if file[rank] {
                count += 1
            }
        }
    }

    return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0

    for _, file := range cb {
        count += len(file)
    }

    return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
    count := 0
    
	for i := 1; i < 9; i+= 1 {
        count += CountInRank(cb, i)
    }

    return count
}
