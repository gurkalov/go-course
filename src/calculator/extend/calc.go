package extend

import "math"

type Calc struct {
	A float64
	B float64
}

func (c Calc) Sum() float64 {
	return c.A + c.B
}

func (c Calc) Sub() float64 {
	return c.A - c.B
}

func (c Calc) Mult() float64 {
	return c.A * c.B
}

func (c Calc) Devide() float64 {
	if c.B == 0 {
		panic("Devide 0")
	}

	return c.A / c.B
}

func (c Calc) Pow() float64 {
	n := uint(c.B)
	res := float64(1)
	for i := uint(0); i < n; i++ {
		res *= c.A
	}

	return res
}

func (c Calc) Mod() float64 {
	return math.Mod(c.A, c.B)
}


