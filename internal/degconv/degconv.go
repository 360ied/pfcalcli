package degconv

import "math"

func DegreesToRadians(deg float64) float64 {
	return deg * math.Pi / 180
}

func RadiansToDegrees(rad float64) float64 {
	return rad * 180 / math.Pi
}
