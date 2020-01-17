package clientapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/Ranadia/overengineered-calculator/calculator"
	googlecloudhandler "github.com/Ranadia/overengineered-calculator/googleCloudHandler"
	"github.com/Ranadia/overengineered-calculator/model"
	"github.com/gorilla/mux"
)

type ClientAPI struct {
	calculator         calculator.Calculator
	googleCloudHandler googlecloudhandler.GoogleCloudHandler
}

func init() {
	fmt.Println("HttpServer starting up.")

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func (ca *ClientAPI) retrieveStaticData(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	data := ca.googleCloudHandler.GetStaticData(ctx)

	jsonData, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err)
	}

	w.Write([]byte(jsonData))
}

func (ca *ClientAPI) receiveCalculation(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	enableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	body, _ := ioutil.ReadAll(r.Body)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	typeOfCalculation := keyVal["typeOfCalculation"]
	fNumber, _ := strconv.ParseFloat(keyVal["firstNumber"], 64)
	sNumber, _ := strconv.ParseFloat(keyVal["secondNumber"], 64)

	var result float64

	switch typeOfCalculation {
	case "plus":
		result = ca.calculator.Plus(fNumber, sNumber)
	case "minus":
		result = ca.calculator.Minus(fNumber, sNumber)
	case "multiply":
		result = ca.calculator.Multiply(fNumber, sNumber)
	case "divide":
		result = ca.calculator.Divide(fNumber, sNumber)
	default:
		w.Write([]byte(`{"message": "Type of calculation is unsupported. Please use either \"plus, minus, multiply, divide\""}`))
		return
	}

	calc := model.Calculation{
		typeOfCalculation,
		fNumber,
		sNumber,
		result,
	}

	ca.googleCloudHandler.PostCalculation(ctx, calc)

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

func (ca *ClientAPI) APIHandle() {
	fmt.Println("apiHandle evoked.")

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/postCalculation", ca.receiveCalculation).Methods(http.MethodPost)

	api.HandleFunc("/staticData", ca.retrieveStaticData).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
}
