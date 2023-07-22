package main

type dirSelect struct {
	action func(dir direction) bool
	
}

func newDirSelect(action func(dir direction) bool) dirSelect {
	return dirSelect{action: action}
}

func (ds dirSelect) processTurn() {
	
}

func (ds dirSelect) processKey(char rune) (validKey bool) {

	if dir, ok := keyToDirMap[char]; ok {
		
		currentState = gameplay
		if ds.action(dir) {
			currentState.processTurn()
		}
		return true
	}

	switch char {
	case 0:
		currentState = gameplay
		gridOverlay.clear()
		return true
	}
	
	return false
}
