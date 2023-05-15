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
	case q:
		os.Exit(0)
	case i:
		showInventory()
	case northKey:
		moveSuccessful = p.move(North)
	case southKey:
		moveSuccessful = p.move(South)
	case eastKey:
		moveSuccessful = p.move(East)
	case westKey:
		moveSuccessful = p.move(West)
	case northWestKey:
		moveSuccessful = p.move(NorthWest)
	case northEastKey:
		moveSuccessful = p.move(NorthEast)
	case southWestKey:
		moveSuccessful = p.move(SouthWest)
	case southEastKey:
		moveSuccessful = p.move(SouthEast)
	case restKey:
		pickUpItem()
		moveSuccessful = p.move(None)
	}
	if moveSuccessful {
		g.processTurn()
		return true
	}
	return false
}