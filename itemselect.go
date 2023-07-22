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
	gridOverlay = nil
	world.Turn++
	lvl.moveMonsters()
}

func (it itemSelect) processKey(char rune) (validKey bool) {

	switch char {
	case 0:
		currentState = gameplay
		gridOverlay = nil
		return true
	case northKey:
		selectedItem--
		if selectedItem < 0 {
			selectedItem = len(menuItems) - 1
		}
	case southKey:
		selectedItem++
		if selectedItem > len(menuItems)-1 {
			selectedItem = 0
		}
	case restKey:
		currentState = itemActions[it.verb](menuItems[selectedItem])
		currentState.processTurn()
		return true
	}
	generateOverlay(true, it.verb)

	return true
}
