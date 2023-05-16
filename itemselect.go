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
	default:
		return false
	}
}