package strain

// Implement the "Keep" and "Discard" function in this file.
func Keep[T any] (arr []T, predicate func(T) bool) []T {
    ret := []T{}

    for _, element := range arr {
        if predicate(element) {
            ret = append(ret, element)
        }
    }

    return ret
}

func Discard[T any] (arr []T, predicate func(T) bool) []T {
    ret := []T{}

    for _, element := range arr {
        if !predicate(element) {
            ret = append(ret, element)
        }
    }

    return ret
}

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1
