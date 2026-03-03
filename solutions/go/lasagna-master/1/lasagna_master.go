package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
    if prepTimePerLayer <= 0 {
        prepTimePerLayer = 2
    }

    return len(layers) * prepTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	totalNoodle := 0
    totalSauce := 0.0
    
    for _, layer := range layers {
        if layer == "sauce" {
            totalSauce += 0.2
        } else if layer == "noodles" {
            totalNoodle += 50
        }
    }

    return totalNoodle, totalSauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
    myList[len(myList) - 1] = friendsList[len(friendsList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, numberOfPortions int) []float64 {
    ret := []float64{}

    for _, amount := range amounts {
        ret = append(ret, (amount / 2) * float64(numberOfPortions))
    }

    return ret
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
