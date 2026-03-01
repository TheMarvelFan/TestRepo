package allergies

import "slices"

func Allergies(allergies uint) []string {
	for allergies > 255 {
        allergies -= 256
    }

    allergenMap := map[int]string{
        1: "eggs",
        2: "peanuts",
        4: "shellfish",
        8: "strawberries",
        16: "tomatoes",
        32: "chocolate",
        64: "pollen",
        128: "cats",
    }

    list := []int{128, 64, 32, 16, 8, 4, 2, 1}
    ret := []string{}

    for _, val := range list {
        if allergies == 0 {
            break
        }

        if uint(val) <= allergies {
            ret = append(ret, allergenMap[val])
            allergies -= uint(val)
        }
    }

    return ret
}

func AllergicTo(allergies uint, allergen string) bool {
	return slices.Contains(Allergies(allergies), allergen)
}
