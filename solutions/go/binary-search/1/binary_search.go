package binarysearch

func SearchInts(list []int, key int) int {
	if len(list) == 0 || key > list[len(list) - 1] {
        return -1
    }
    
	lo := 0
    hi := len(list)

    for lo <= hi {
        mid := lo + (hi - lo) / 2

        if list[mid] == key {
            return mid
        }

        if list[mid] < key {
            lo = mid + 1
        } else {
            hi = mid - 1
        }
    }

    return -1
}
