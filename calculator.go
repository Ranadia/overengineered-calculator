package main

import (
	"context"
)

func plus(a float64, b float64) float64 {
	result := a + b
	ctx := context.Background()

	calc := Calculation{"plus", a, b, float64(result)}

	postCalculation(ctx, calc)
	return result
}

func minus(a float64, b float64) float64 {
	result := a - b
	ctx := context.Background()

	calc := Calculation{"minus", a, b, result}

	postCalculation(ctx, calc)
	return result
}

func multiply(a float64, b float64) float64 {
	result := a * b
	ctx := context.Background()

	calc := Calculation{"multiply", a, b, result}

	postCalculation(ctx, calc)
	return result
}

func divide(a float64, b float64) float64 {
	result := a / b
	ctx := context.Background()

	calc := Calculation{"divide", a, b, result}

	postCalculation(ctx, calc)
	return result
}
