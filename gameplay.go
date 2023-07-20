package main

import (
	"fmt"
	"os"
)

type gamePlay struct {
}

func (g gamePlay) processTurn() {
	gridOverlay = nil
	turn++
	moveMonsters()
}

func (g gamePlay) processKey(char rune) (validKey bool) {
	fmt.Println(char)
	var moveSuccessful bool = false

	if dir, ok := keyToDirMap[char]; ok {
		moveSuccessful = p.attemptMove(dir)
	}

	switch char {
	case quitKey:
		persistance.saveState("save.sav")
		os.Exit(0)
	case openKey:
		open()
		return true
	case closeKey:
		close()
		return true
	case lookKey:
		look()
		return true
	case inventoryKey:
		generateOverlay(false, "drop")
		return true
	case dropKey:
		currentState = newItemSelect("drop")
		return true
	case eatKey:
		currentState = newItemSelect("eat")
		return true
	case throwKey:
		currentState = newItemSelect("throw")
		return true
	case restKey:
		p.pickUpItem()
		return true
	case downStairs:
		useStairs("down")
		return true
	case upStairs:
		useStairs("up")
		return true
	}
	if moveSuccessful {
		checkPosition()
		g.processTurn()
		return true
	}
	return false
}
