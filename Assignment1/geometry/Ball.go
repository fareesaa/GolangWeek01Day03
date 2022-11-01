package geometry

import (
	"math"
)

type Ball struct {
	Radius float64
}

func (ball Ball) VolBall() float64 {
	return (4 * math.Pi * ball.Radius * ball.Radius * ball.Radius) / 3
}
