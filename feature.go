package main

type feature struct {
	position    point
	prefix      string
	name        string
	char        string
	description string
	closed      bool
}

func createDoor(position point) (*feature, bool) {

	if position.isInCorridor() {
		door := feature{
			position: position,
			prefix:   "a",
			name:     "door",
		}
		if randomNumber(2) == 1 {
			door.closed = true
			door.char = "+"
			door.description = "a closed door"
		} else {
			door.closed = false
			door.char = "-"
			door.description = "an open door"
		}
		return &door, true
	} else {
		return &feature{}, false
	}

}

func (f feature) getChar() rune {
	return rune(f.char[0])
}
