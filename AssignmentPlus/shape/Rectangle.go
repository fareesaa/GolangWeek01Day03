package shape

import ()

type Rectangle struct {
	S int
}

func (rectangle Rectangle) AreaRect() int {
	return rectangle.S * rectangle.S
}
