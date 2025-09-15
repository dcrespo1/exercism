// Package weather is used to ditermine the forcast for a given location.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location of the forcast.
var CurrentLocation string

// Forecast takes a city and a condition string and returns a string that containes
// the CurrentLocation and the CurrentCondition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
