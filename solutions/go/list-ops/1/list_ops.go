package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	res := initial

    for _, num := range s {
        res = fn(res, num)
    }

    return res
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	res := initial

    for i := s.Length() - 1; i > -1; i -- {
        res = fn(s[i], res)
    }

    return res
}

func (s IntList) Filter(fn func(int) bool) IntList {
	res := make([]int, s.Length())
    idx := 0

    for _, num := range s {
        if fn(num) {
        	res[idx] = num
            idx ++
        }
    }

    return IntList(res[:idx])
}

func (s IntList) Length() int {
	count := 0

    for _, num := range s {
        count += num - num + 1
    }

    return count
}

func (s IntList) Map(fn func(int) int) IntList {
	res := make([]int, s.Length())

    for idx, num := range s {
        res[idx] = fn(num)
    }

    return IntList(res)
}

func (s IntList) Reverse() IntList {
	res := make([]int, s.Length())
    idx := 0

    for i := s.Length() - 1; i > -1; i -- {
        res[idx] = s[i]
        idx ++
    }

    return IntList(res)
}

func (s IntList) Append(lst IntList) IntList {
	res := make([]int, s.Length() + lst.Length())
    idx := 0

    for _, num := range s {
        res[idx] = num
        idx ++
    }

    for _, num := range lst {
        res[idx] = num
        idx ++
    }

    return IntList(res)
}

func (s IntList) Concat(lists []IntList) IntList {
	res := s

    for _, list := range lists {
        res = res.Append(list)
    }

    return res
}
