package simple

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
