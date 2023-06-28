package main

type dirSelect struct {
	action func(dir direction)
	
}

func newDirSelect(action func(dir direction)) dirSelect {
	return dirSelect{action: action}
}

func (ds dirSelect) processTurn() {
	
}

func (ds dirSelect) processKey(char rune) (validKey bool) {

	if dir, ok := keyToDirMap[char]; ok {
		ds.action(dir)
		currentState = gameplay
		return true
	}

	switch char {
	case 0:
		currentState = gameplay
		gridOverlay = nil
		return true
	}
	//default: return false?
	return false
}
