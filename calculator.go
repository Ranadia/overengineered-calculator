package main

func plus(a int, b int) int {
	result := a + b
	postCalculation("plus", a, b, result)
	return result
}

func minus(a int, b int) int {
	result := a - b
	postCalculation("minus", a, b, result)
	return result
}

func multiply(a int, b int) int {
	result := a * b
	postCalculation("multiply", a, b, result)
	return result
}

func divide(a int, b int) int {
	result := a / b
	postCalculation("divide", a, b, result)
	return result
}
