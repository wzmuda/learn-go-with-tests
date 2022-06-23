package main

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}
