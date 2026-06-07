package structs

import (
	"math"
	"testing"
)

func TestCircle(t *testing.T) {

	circleParams := CircleParams{radius: 3}

	got := CircleArea(circleParams)

	want := math.Pi * math.Pow(circleParams.radius, 2)

	//fmt.Printf("got %g, want %g", got, want)

	if got != want {
		t.Errorf("got %.5f, want %.5f", got, want)
	}
}
