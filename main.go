package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Printf("Initialized calculator\n")
	fmt.Printf("Doing Calculations\n")
	plus(3.5, 1.4)
	minus(20.8, 17.9)
	multiply(8.7, 26.3)
	divide(297.3, 13.2)

	fmt.Println("\nGetting static data")
	getStaticData(ctx)

	fmt.Println("\nGetting data for plus calculations")
	getCalculation(ctx, "plus")

	closeClient()
}
