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

	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{5.0, 5.0}
		got := rectangle.Area()
		want := 25.0

		if got != want {
			t.Errorf("Got %.2f, want %.2f", got, want)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{5.0}
		got := circle.Area()
		want := 78.53981633974483

		if got != want {
			t.Errorf("Got %g, want %g", got, want)
		}

	})

}
