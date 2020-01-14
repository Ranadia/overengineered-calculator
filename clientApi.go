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

type server struct{}

func init() {
	fmt.Println("HttpServer starting up.")
}

func retrieveAllCalculations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "get called"}`))
}

func retrieveCalculationForType(w http.ResponseWriter, r *http.Request) {
	pathParams := mux.Vars(r)
	ctx := context.Background()

	typeOfCalculation := ""
	var err error
	if val, ok := pathParams["typeOfCalculation"]; ok {
		typeOfCalculation = val
		if err != nil {
			fmt.Println("Need a typeOfCalculation of either plus, minus, multiply or divide")
		}
	}

	if typeOfCalculation == "plus" || typeOfCalculation == "minus" || typeOfCalculation == "multiply" || typeOfCalculation == "divide" {
		getCalculationsForType(ctx, typeOfCalculation)
	} else {
		fmt.Println("Type of calculation not recognised. Please use either \"plus, minus, multiply or divide\"")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func retrieveStaticData(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	data := getStaticData(ctx)

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(jsonData))
}

func receiveCalculation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body) // check for errors

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal) // check for errors

	typeOfCalculation := keyVal["typeOfCalculation"]
	fNumber, _ := strconv.ParseFloat(keyVal["firstNumber"], 64)
	sNumber, _ := strconv.ParseFloat(keyVal["secondNumber"], 64)

	var result float64

	switch typeOfCalculation {
	case "plus":
		result = plus(fNumber, sNumber)
	case "minus":
		result = minus(fNumber, sNumber)
	case "multiply":
		result = multiply(fNumber, sNumber)
	case "divide":
		result = divide(fNumber, sNumber)
	default:
		w.Write([]byte(`{"message": "Type of calculation is unsupported. Please use either \"plus, minus, multiply, divide\""}`))
		return
	}

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

func apiHandle() {
	fmt.Println("apiHandle evoked.")

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/getAll", retrieveAllCalculations).Methods(http.MethodGet)
	api.HandleFunc("/getCalculation/{typeOfCalculation}", retrieveCalculationForType).Methods(http.MethodGet)

	api.HandleFunc("/postCalculation", receiveCalculation).Methods(http.MethodPost)

	api.HandleFunc("/staticData", retrieveStaticData).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
