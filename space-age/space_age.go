package space

var planet_data = map[string]float64{
	"Earth":   31557600.0,
	"Mercury": 31557600.0 * 0.2408467,
	"Venus":   31557600.0 * 0.61519726,
	"Mars":    31557600.0 * 1.8808158,
	"Jupiter": 31557600.0 * 11.862615,
	"Saturn":  31557600.0 * 29.447498,
	"Uranus":  31557600.0 * 84.016846,
	"Neptune": 31557600.0 * 164.79132,
}

func Age(secs float64, planetName Planet) float64 {
	name := string(planetName)
	time := float64(secs)
	result := time / planet_data[name]
	return result
}
