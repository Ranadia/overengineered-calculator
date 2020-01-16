package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ClientAPI struct {
	calculator *Calculator
	gch        *googleCloudHandler
}

func init() {
	fmt.Println("HttpServer starting up.")
}

func (ca ClientAPI) retrieveStaticData(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data := ca.gch.getStaticData(ctx)

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(jsonData))
}

func (ca ClientAPI) receiveCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.Background()

	body, _ := ioutil.ReadAll(r.Body)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	typeOfCalculation := keyVal["typeOfCalculation"]
	fNumber, _ := strconv.ParseFloat(keyVal["firstNumber"], 64)
	sNumber, _ := strconv.ParseFloat(keyVal["secondNumber"], 64)

	var result float64

	switch typeOfCalculation {
	case "plus":
		result = ca.calculator.plus(fNumber, sNumber)
	case "minus":
		result = ca.calculator.minus(fNumber, sNumber)
	case "multiply":
		result = ca.calculator.multiply(fNumber, sNumber)
	case "divide":
		result = ca.calculator.divide(fNumber, sNumber)
	default:
		w.Write([]byte(`{"message": "Type of calculation is unsupported. Please use either \"plus, minus, multiply, divide\""}`))
		return
	}

	calc := Calculation{
		typeOfCalculation,
		fNumber,
		sNumber,
		result,
	}

	ca.gch.postCalculation(ctx, calc)

	data := map[string]interface{}{
		"result": result,
	}
	jsonData, err := json.Marshal(data)

	if err != nil {
		w.Write([]byte(`{"message": "Error json.marshalling result."}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonData))
}

func (ca *ClientAPI) apiHandle() {
	fmt.Println("apiHandle evoked.")

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	// api.HandleFunc("/getAll", retrieveAllCalculations).Methods(http.MethodGet)
	// api.HandleFunc("/getCalculation/{typeOfCalculation}", retrieveCalculationForType).Methods(http.MethodGet)

	api.HandleFunc("/postCalculation", ca.receiveCalculation).Methods(http.MethodPost)

	api.HandleFunc("/staticData", ca.retrieveStaticData).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
