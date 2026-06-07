package structs

import "testing"

func TestPerimeter(t *testing.T) {

	square := Square{height: 10, width: 10}

	//given h, w = 10
	got := Perimeter(square)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}

}
