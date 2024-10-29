package sprint

import "math"

func Mean(a, b, c float32) float32 {
	mean := (a + b + c) / 3
	rounded := float32(math.Round(float64(mean)*10) / 10)
	return rounded
}
