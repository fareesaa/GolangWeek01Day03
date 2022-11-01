package geometry

import ()

type Cube struct {
	S int
}

func (cube Cube) VolCube() int {
	return cube.S * cube.S * cube.S
}
