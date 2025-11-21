package dndcharacter

import "math/rand"
import "math"

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return int(math.Floor(float64(score - 10) / 2.0))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	scores := []int{0, 0, 0}
    
    for i := 0; i < 4; i ++ {
		n := rand.Intn(5) + 1
        
        if n > scores[0] {
            scores[0] = n
        } else if n > scores[1] {
            scores[1] = n
        } else if n > scores[2] {
            scores[2] = n
        }
    }

    sum := 0

    for _, score := range scores {
        sum += score
    }

    return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	ch := Character{
        Strength: Ability(),
        Dexterity: Ability(),
        Constitution: Ability(),
        Intelligence: Ability(),
        Wisdom: Ability(),
        Charisma: Ability(),
    }

    ch.Hitpoints = 10 + Modifier(ch.Constitution)

    return ch
}
