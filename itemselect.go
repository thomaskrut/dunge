package main

type itemSelect struct {
	verb string
}

func newItemSelect(v string) itemSelect {
	return itemSelect{verb: v}
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
		if inventoryMenu.getItemByNumber(int(index)) != nil {
			itemAction(it.verb, inventoryMenu.getItemByNumber(int(index)))
			currentState = gameplay
			return true
		}

	}
	return false
}