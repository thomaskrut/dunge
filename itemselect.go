package main

type itemSelect struct {
	verb string
}

func newItemSelect(v string) itemSelect {
	gridOverlay.selection = 0
	gridOverlay.generate(true, v)
	return itemSelect{verb: v}
}

func (it itemSelect) processTurn() {
	gridOverlay.clear()
	world.Turn++
	lvl.moveMonsters()
}

func (it itemSelect) processKey(char rune) (validKey bool) {

	switch char {
	case 0:
		currentState = gameplay
		gridOverlay.clear()
		return true
	case northKey:
		gridOverlay.cursorUp()
	case southKey:
		gridOverlay.cursorDown()
	case restKey:
		currentState = itemActions[it.verb](gridOverlay.selectedItem())
		currentState.processTurn()
		return true
	}
	gridOverlay.generate(true, it.verb)

	return true
}
