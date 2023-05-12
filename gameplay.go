package main

import (

	"os"
)
type gamePlay struct {
}

func (g gamePlay) processTurn() {
	moveMonsters()
	checkMonsterHealth()
}

func (g gamePlay) processKey(char rune) bool {
	
	var playerMoved bool = false
	switch char {
		case q:
			os.Exit(0)
	case northKey:
		playerMoved = p.move(North)
	case southKey:
		playerMoved = p.move(South)
	case eastKey:
		playerMoved = p.move(East)
	case westKey:
		playerMoved = p.move(West)
	case northWestKey:
		playerMoved = p.move(NorthWest)
	case northEastKey:
		playerMoved = p.move(NorthEast)
	case southWestKey:
		playerMoved = p.move(SouthWest)
	case southEastKey:
		playerMoved = p.move(SouthEast)
	case restKey:
		playerMoved = p.move(None)
	}
	if playerMoved {
		g.processTurn()
		return true
	}
	return false
}