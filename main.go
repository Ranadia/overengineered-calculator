package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Initialized calculator\n")

	fmt.Println(plus(1, 2))
	fmt.Println(minus(10, 4))
	fmt.Println(multiply(10, 10))
	fmt.Println(divide(20, 4))
}
