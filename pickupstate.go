package main

import (
	"fmt"
	"os"
)

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
	switch char {
	case q:
		os.Exit(0) 
	case notAChar :
		currentState = gamePlay{}
		return true
	}
	
	return false
}