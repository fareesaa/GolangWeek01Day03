package shape

import ()

type Triangle struct {
	A, T int
}

func (triangle Triangle) AreaTri() int {
	return (triangle.A * triangle.T) / 2
}
