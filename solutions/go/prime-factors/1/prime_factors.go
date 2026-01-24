package prime

func Factors(n int64) []int64 {
	ret := []int64{}

    for n % 2 == 0 {
        ret = append(ret, 2)
        n /= 2
    }

    var i int64

    for i = 3; i <= n; i += 2 {
        for n % i == 0 {
            ret = append(ret, i)
            n /= i
        }
    }

    if n > 2 {
        ret = append(ret, n)
    }

    return ret
}
