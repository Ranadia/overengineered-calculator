package main

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
)

func init() {
	const projectID = ""

	ctx := context.Background()
	conf := &firebase.Config{ProjectID: projectID}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		fmt.Println("App not initialized\n")
		fmt.Println(err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("Client not initialized\n")
		fmt.Println(err)
	}
}

func postCalculation(typeOfCalc string, firstInteger int, secondInteger int, result int) {
	fmt.Println(typeOfCalc)
	fmt.Println(firstInteger)
	fmt.Println(secondInteger)
	fmt.Println(result)

	client.Collection("calculations").Add(ctx, map[string]interface{}{
		"typeOfCalculation": typeOfCalc,
		"firstInteger":      firstInteger,
		"secondInteger":     secondInteger,
		"result":            result,
	})
}

func getCalculation() {

}
