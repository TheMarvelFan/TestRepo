package pascalstriangle

func Triangle(n int) [][]int {
	if n <= 0 {
        return nil
    }
    
    var row []int
	ret := [][]int{{1}}
    
    for i := 1; i < n; i++ {
        row = []int{1}
        
        for j := 0; j < len(ret[i - 1]) - 1; j++ {
            row = append(row, ret[i - 1][j] + ret[i - 1][j + 1])
        }
        
        row = append(row, 1)
        ret = append(ret, row)
    }
    
    return ret
}
