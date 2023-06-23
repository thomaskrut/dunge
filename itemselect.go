package main

type itemSelect struct {
	verb string
}

func newItemSelect(v string) itemSelect {
	selectedItem = 0
	generateOverlay(true, v)
	return itemSelect{verb: v}
}

func (it itemSelect) processTurn() {

}

func (it itemSelect) processKey(char rune) bool {

	switch char {
	case 0:
		currentState = gameplay
		previousState = currentState
	case northKey:
		selectedItem--
	case southKey:
		selectedItem++
	case restKey:
		itemAction(it.verb)
	}
	generateOverlay(true, it.verb)

	return true
}
