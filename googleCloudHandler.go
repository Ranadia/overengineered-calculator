package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
)

var (
	client    *firestore.Client
	projectID string
)

type googleCloudHandler struct {
}

func init() {
	projectID = "overengineered-calculato-c695d"
	ctx := context.Background()
	var err error

	client, err = firestore.NewClient(ctx, projectID)

	if err != nil {
		fmt.Println(err)
	}
}

func (gch *googleCloudHandler) postCalculation(ctx context.Context, calc Calculation) {
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

func (gch *googleCloudHandler) getStaticData(ctx context.Context) map[string]interface{} {
	staticData, err := client.Collection("staticData").Doc("AppData").Get(ctx)

	if err != nil {
		fmt.Println(err)
	}

	data := staticData.Data()

	return data
}

func (gch *googleCloudHandler) closeClient() {
	fmt.Println("Closing client")
	client.Close()
	fmt.Println("Client closed")
}
