package squareroot

func SquareRoot(x int) (int, error) {
	r := x;
    
    for r * r > x {
        r = (r + x/r) / 2;
    }
    
    return r, nil;
}
