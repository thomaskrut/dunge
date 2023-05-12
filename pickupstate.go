package main

import (
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
	
	switch char {
	case q:
		os.Exit(0) 
	case notAChar :
		currentState = gamePlay{}
		return true
	}

	digit := int(convertToDigit(char))

	if digit > 0 && digit <= p.numberOfItems {
		pickUpItem(digit)
		currentState = gamePlay{}
		return true
	}
	
	return false
}