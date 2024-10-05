package structs

type Rectangle struct {
	width  float64
	height float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.width + r.height)
}

func Area(width, height float64) float64 {
	return width * height

}
