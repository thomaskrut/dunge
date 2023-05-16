package main

import (
	"fmt"
	"os"
)

type gamePlay struct {
}

func (g gamePlay) processTurn() {
	moveMonsters()
	checkForItems()
}

func (g gamePlay) processKey(char rune) bool {
	fmt.Println(char)
	var moveSuccessful bool = false
	switch char {
	case quitKey:
		os.Exit(0)
	case inventoryKey:
		printInventory("Inventory", "all")
	case dropKey:
		currentState = newItemSelect("drop")
		return true
	case northKey:
		moveSuccessful = p.attemptMove(North)
	case southKey:
		moveSuccessful = p.attemptMove(South)
	case eastKey:
		moveSuccessful = p.attemptMove(East)
	case westKey:
		moveSuccessful = p.attemptMove(West)
	case northWestKey:
		moveSuccessful = p.attemptMove(NorthWest)
	case northEastKey:
		moveSuccessful = p.attemptMove(NorthEast)
	case southWestKey:
		moveSuccessful = p.attemptMove(SouthWest)
	case southEastKey:
		moveSuccessful = p.attemptMove(SouthEast)
	case restKey:
		pickUpItem()
		moveSuccessful = p.attemptMove(None)
	}
	if moveSuccessful {
		g.processTurn()
		return true
	}
	return false
}
