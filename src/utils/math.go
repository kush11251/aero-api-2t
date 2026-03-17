package utils

func CalculateWindSpeed(windDirection float64, windForce float64) float64 {
	return windForce * windDirection
}