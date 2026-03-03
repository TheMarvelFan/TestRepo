// Package weather provides tools to correctly forcast the weather at Goblinocus.
package weather

var (
    // CurrentCondition provides the current weather condition for the specified city.
	CurrentCondition string
    // CurrentLocation specifies the city for which the weather is being predicted.
	CurrentLocation  string
)

// Forecast actually forecasts the weather for the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
