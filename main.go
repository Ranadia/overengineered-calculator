package main

import (
	clientapi "github.com/Ranadia/overengineered-calculator/clientAPI"
)

func main() {
	clientAPI := &clientapi.ClientAPI{}
	clientAPI.APIHandle()
}
