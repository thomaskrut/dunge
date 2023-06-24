package main

type itemSelect struct {
	verb string
}

func newItemSelect(v string) itemSelect {
	generateOverlay(true, v)
	return itemSelect{verb: v}
}

func (it itemSelect) processTurn() {
	gridOverlay = nil
	turn++
	moveMonsters()
}

func (it itemSelect) processKey(char rune) bool {

	switch char {
	case 0:
		currentState = gameplay
	case northKey:
		selectedItem--
	case southKey:
		selectedItem++
	case restKey:
		itemAction(it.verb)
		currentState.processTurn()
		currentState = gameplay
		return true
	}
	generateOverlay(true, it.verb)

	return true
}
