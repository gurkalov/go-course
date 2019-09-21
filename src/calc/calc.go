package rpn

func Calc (op string, values ...float64) float64 {
	result := values[0]
	for i := range values[1:] {
		switch op {
		case "+":
			result += values[i + 1]
		case "-":
			result -= values[i + 1]
		case "*":
			result *= values[i + 1]
		case "/":
			result /= values[i + 1]
		default:
			panic("Error operator")
		}
	}

	return result
}