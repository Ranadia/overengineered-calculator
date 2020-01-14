package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	ctx := context.Background()

	fmt.Println("\nprinting r:")
	fmt.Println(r)
	fmt.Println("\n\n")

	err := r.ParseForm()

	fNumber, _ := strconv.ParseFloat(r.FormValue("firstNumber"), 64)
	sNumber, _ := strconv.ParseFloat(r.FormValue("secondNumber"), 64)
	rslt, _ := strconv.ParseFloat(r.FormValue("result"), 64)

	fmt.Println("\nreceiveCalculation, printing r:")
	fmt.Println("typeOfCalculation: ", r.FormValue("typeOfCalculation"))
	fmt.Println("firstNumber: ", r.FormValue("firstNumber"))
	fmt.Println("secondNumber: ", r.FormValue("secondNumber"))
	fmt.Println("result: ", r.FormValue("result"))
	fmt.Println("\n\n")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}

	calc := Calculation{
		typeOfCalculation: r.FormValue("typeOfCalculation"),
		firstNumber:       fNumber,
		secondNumber:      sNumber,
		result:            rslt,
	}

	postCalculation(ctx, calc)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message": "receiveCalculation test"}`))
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
