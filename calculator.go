package main

type Calculator struct{}

func (c *Calculator) plus(a float64, b float64) float64 {
	return a + b
}

func (c *Calculator) minus(a float64, b float64) float64 {
	return a - b
}

func (c *Calculator) multiply(a float64, b float64) float64 {
	return a * b
}

// TODO: Implement handling for dividing by zero
func (c *Calculator) divide(a float64, b float64) float64 {
	return a / b
}
