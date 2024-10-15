package structs

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{Width: 5.0, Height: 5.0}, 25.0},
		{"Circle", Circle{Radius: 5.0}, 78.53981633974483},
		{"Triangle", Triangle{Base: 5, Height: 2}, 5.0},
	}

	for _, at := range areaTests {
		t.Run(at.name, func(t *testing.T) {
			got := at.shape.Area()
			if got != at.want {
				t.Errorf("%#v Got %.2f, want %.2f", at.shape, got, at.want)
			}
		})
	}

}
