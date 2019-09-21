package calc

import "course/calc/extend"

type Pair struct {
	A float64
	B float64
}

type Calc interface {
	Sum() float64
	Sub() float64
	Mult() float64
	Devide() float64
}

func Compute (c Calc, op string) float64 {
	switch op {
		case "+": {
			return c.Sum()
		}
		case "-": {
			return c.Sub()
		}
		case "*": {
			return c.Mult()
		}
		case "/": {
			return c.Devide()
		}
	}

	panic("Invalid operator")
	return 0
}

func ComputeExtend (c extend.Calc, op string) float64 {
	switch op {
	case "^": {
		return c.Pow()
	}
	case "mod": {
		return c.Mod()
	}
	}

	return Compute(c, op)
}
