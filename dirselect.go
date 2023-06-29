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
		if actionSuccessful := ds.action(dir); actionSuccessful {
			currentState.processTurn()
		}
		return true
	}

	switch char {
	case 0:
		currentState = gameplay
		gridOverlay = nil
		return true
	}
	
	return false
}
