package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    total := 0
    
	for _, val := range birdsPerDay {
        total += val
    }

    return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsPerWeek := []int{}
    total := 0

    for id, val := range birdsPerDay {
        if id % 7 == 0 {
            birdsPerWeek = append(birdsPerWeek, total)
            total = 0
        }
        total += val
    }

    if (total > 0) {
        birdsPerWeek = append(birdsPerWeek, total)
        total = 0
    }

    return birdsPerWeek[week]
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for id, _ := range birdsPerDay {
        if id % 2 == 0 {
            birdsPerDay[id] += 1
        }
    }

    return birdsPerDay
}
