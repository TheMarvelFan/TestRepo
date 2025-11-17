package space

import "strings"

type Planet string

func Age(seconds float64, planet Planet) float64 {
	ageInSecondsOnPlanet := map[Planet]float64{}

    ageInSecondsOnPlanet["mercury"] = 0.2408467 * 31_557_600.0
    ageInSecondsOnPlanet["venus"] = 0.61519726 * 31_557_600.0
    ageInSecondsOnPlanet["earth"] = 31_557_600.0
    ageInSecondsOnPlanet["mars"] = 1.8808158 * 31_557_600.0
    ageInSecondsOnPlanet["jupiter"] = 11.862615 * 31_557_600.0
    ageInSecondsOnPlanet["saturn"] = 29.447498 * 31_557_600.0
    ageInSecondsOnPlanet["uranus"] = 84.016846 * 31_557_600.0
    ageInSecondsOnPlanet["neptune"] = 164.79132 * 31_557_600.0

    val, ok := ageInSecondsOnPlanet[Planet(strings.ToLower(string(planet)))]

    if !ok {
        return -1.0
    }

    return seconds / val
}
