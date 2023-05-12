package main

import "fmt"

type pickupState struct {
	numberOfItems int
}

func (p pickupState) processTurn() {
	
}

func newPickupState(count int) pickupState {
	return pickupState{numberOfItems: count}
}

func (p pickupState) processKey(char rune) bool {
	fmt.Println(char)
	currentState = gamePlay{}
	return true
}