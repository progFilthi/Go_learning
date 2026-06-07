package structs

type Rectangle struct {
	length float64
	width  float64
}

func Area(rectangle Rectangle) float64 {

	return rectangle.length * rectangle.width
}
