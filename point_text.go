package sprint

import "fmt"

type Point struct {
	X    float64
	Y    float64
	Text string
}

func PointText(p Point) Point {
	return Point{
		X:    p.X,
		Y:    p.Y,
		Text: fmt.Sprintf("Text at (%.6f, %.6f)", p.X, p.Y),
	}
}
