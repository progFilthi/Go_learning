package structs

type Square struct {
	height float64
	width  float64
}

func Perimeter(square Square) float64 {

	perimeter := 2 * (square.height + square.width)

	return perimeter
}
