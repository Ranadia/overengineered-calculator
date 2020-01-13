package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

var client *firestore.Client

var projectID string

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
	// client, err := firestore.NewClient(ctx, projectID)

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

func getCalculation(ctx context.Context, typeOfCalculation string) {
	// client, err := firestore.NewClient(ctx, projectID)

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

func getStaticData(ctx context.Context) {
	staticData, err := client.Collection("staticData").Doc("AppData").Get(ctx)

	if err != nil {
		fmt.Println(err)
	}

	data := staticData.Data()

	fmt.Println(data["Name"])
}

func closeClient() {
	fmt.Println("Closing client")
	client.Close()
	fmt.Println("Client closed")
}
