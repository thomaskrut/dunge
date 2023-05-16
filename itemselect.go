package main

type itemSelect struct {
	verb string
	currentMenu map[int]string
}

func newItemSelect(v string) itemSelect {
	menu := generateInventory()
	return itemSelect{verb: v, currentMenu: menu}
}

func (it itemSelect) processTurn() {

}

func (it itemSelect) processKey(char rune) bool {

	switch char {
	case 0, dropKey:
		currentState = gameplay
		return true
	}

	if char > 48 && char < 58 {
		index := convertToDigit(char)
		if _, ok := it.currentMenu[int(index)]; ok {
			itemAction(it.verb, it.currentMenu[int(index)])
			return true
		}
	}
	return false
}