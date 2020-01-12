package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

var client firestore

func init() {
	projectID := "overengineered-calculato-c695d"

	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)

	if err != nil {
		fmt.Println(err)
	}

}

func postCalculation(ctx context.Context, calc Calculation) {
	fmt.Println(calc.typeOfCalculation)
	fmt.Println(calc.firstNumber)
	fmt.Println(calc.secondNumber)
	fmt.Println(calc.result)

	_, _, err := client.Collection("calculations").Add(ctx, map[string]interface{}{
		"typeOfCalculation": string(calc.typeOfCalculation),
		"firstInteger":      string(calc.firstNumber),
		"secondInteger":     string(calc.secondNumber),
		"result":            string(calc.result),
	})
	if err != nil {
		fmt.Println(err)
	}
}

func getCalculation() {

}
