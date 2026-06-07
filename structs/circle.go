package structs

import "math"

type CircleParams struct {
	radius float64
}

func CircleArea(circleParams CircleParams) float64 {

	return math.Pi * math.Pow(circleParams.radius, 2)

}
