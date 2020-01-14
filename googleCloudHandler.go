package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

var (
	client    *firestore.Client
	projectID string
)

func init() {
	projectID = "overengineered-calculato-c695d"
	ctx := context.Background()
	var err error

	client, err = firestore.NewClient(ctx, projectID)

	if err != nil {
		fmt.Println(err)
	}
}

func postCalculation(ctx context.Context, calc Calculation) {
	_, _, err := client.Collection(calc.typeOfCalculation).Add(ctx, map[string]interface{}{
		"firstNumber":  calc.firstNumber,
		"secondNUmber": calc.secondNumber,
		"result":       calc.result,
	})

	if err != nil {
		fmt.Println("error with calculationDoc.Set()")
		fmt.Println(err)
	}
}

func getCalculationsForType(ctx context.Context, typeOfCalculation string) {
	iter := client.Collection(typeOfCalculation).Documents(ctx)

	for {
		calc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate: %v", err)
		}

		fmt.Println(calc.Data())
	}
}

func getStaticData(ctx context.Context) map[string]interface{} {
	staticData, err := client.Collection("staticData").Doc("AppData").Get(ctx)

	if err != nil {
		fmt.Println(err)
	}

	data := staticData.Data()

	return data
}

func closeClient() {
	fmt.Println("Closing client")
	client.Close()
	fmt.Println("Client closed")
}
