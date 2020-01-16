package googlecloudhandler

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/Ranadia/overengineered-calculator/model"
)

var (
	client    *firestore.Client
	projectID string
)

type GoogleCloudHandler struct{}

func init() {
	projectID = "overengineered-calculato-c695d"
	ctx := context.Background()
	var err error

	client, err = firestore.NewClient(ctx, projectID)

	if err != nil {
		fmt.Println(err)
	}
}

func (gch *GoogleCloudHandler) PostCalculation(ctx context.Context, calc model.Calculation) {
	_, _, err := client.Collection(calc.TypeOfCalculation).Add(ctx, map[string]interface{}{
		"firstNumber":  calc.FirstNumber,
		"secondNUmber": calc.SecondNumber,
		"result":       calc.Result,
	})

	if err != nil {
		fmt.Println("error with Collection.Add")
		fmt.Println(err)
	}
}

func (gch *GoogleCloudHandler) GetStaticData(ctx context.Context) map[string]interface{} {
	staticData, err := client.Collection("staticData").Doc("AppData").Get(ctx)

	if err != nil {
		fmt.Println(err)
	}

	data := staticData.Data()

	return data
}

func (gch *GoogleCloudHandler) CloseClient() {
	fmt.Println("Closing client")
	client.Close()
	fmt.Println("Client closed")
}
