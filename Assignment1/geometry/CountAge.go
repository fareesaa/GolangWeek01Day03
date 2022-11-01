package geometry

import ()

type CountAge struct {
	B, L int
}

func (countAge CountAge) CountingAge() int {
	return countAge.L - countAge.B
}
