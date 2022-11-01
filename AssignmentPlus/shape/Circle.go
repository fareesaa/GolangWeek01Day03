package shape

import (
	"math"
)

type Circle struct {
	Radius float64
}

func (circle Circle) AreaCir() float64 {
	return math.Pi * circle.Radius * circle.Radius
}
