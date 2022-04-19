// Package weather provides the ability to report weather conditions of locations.
package weather
// CurrentCondition represents the current weather condition of a Location.
var CurrentCondition string
// CurrentLocation represents the location of CurrentCondition.
var CurrentLocation string
// Forecast returns a string value which reports the current weather condition of a location from the given city and condition string input variables.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
