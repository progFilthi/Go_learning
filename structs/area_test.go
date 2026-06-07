package structs

import "testing"

func TestArea(t *testing.T) {

	rectangle := Rectangle{4, 3}

	got := Area(rectangle)
	want := 12.0

	if got != want {
		t.Errorf("got %.2f, want %.2f", got, want)
	}
}
