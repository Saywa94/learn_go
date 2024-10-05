package structs

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	got := Area(5.0, 5.0)
	want := 25.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}

}
