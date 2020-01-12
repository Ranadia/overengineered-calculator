package main

import (
	"context"
	"fmt"
)

func init() {

}

func postCalculation(ctx context.Context, calc Calculation) {
	fmt.Println(calc.typeOfCalculation)
	fmt.Println(calc.firstInteger)
	fmt.Println(calc.secondInteger)
	fmt.Println(calc.result)

}

func getCalculation() {

}
