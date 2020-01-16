package calculator

type Calculator struct{}

func (c *Calculator) Plus(a float64, b float64) float64 {
	return a + b
}

func (c *Calculator) Minus(a float64, b float64) float64 {
	return a - b
}

func (c *Calculator) Multiply(a float64, b float64) float64 {
	return a * b
}

// TODO: Implement handling for dividing by zero
func (c *Calculator) Divide(a float64, b float64) float64 {
	return a / b
}
