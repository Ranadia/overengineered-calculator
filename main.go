package main

import (
	clientapi "github.com/Ranadia/overengineered-calculator/clientAPI"
)

var (
	clientAPI *clientapi.ClientAPI
)

func main() {
	clientAPI.APIHandle()
}
