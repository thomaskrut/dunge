package main

type feature struct {
	Position    point
	Prefix      string
	Name        string
	Char        string
	Description string
	State       string
}

func createDoor(position point) (*feature, bool) {

	if position.isInCorridor() {
		door := feature{
			Position: position,
			Prefix:   "a",
			Name:     "door",
		}
		if randomNumber(2) == 1 {
			door.State = "closed"
			door.Char = "+"
			door.Description = "a closed door"
		} else {
			door.State = "open"
			door.Char = "-"
			door.Description = "an open door"
		}
		return &door, true
	} else {
		return nil, false
	}

}

func createStairs(position point, direction string) (*feature, bool) {
	if _, ok := lvl.Features[position]; ok || position.isInCorridor() {
		return nil, false
	}

	stairs := feature{
		Position: position,
		Name:     "staircase",
		Prefix:   "a",
		State:    direction,
	}

	switch direction {
	case "up":
		stairs.Char = "<"
		stairs.Description = "a staircase going up"
	case "down":
		stairs.Char = ">"
		stairs.Description = "a staircase going down"
	default:
		return nil, false
	}

	return &stairs, true

}

func (f feature) getChar() rune {
	return rune(f.Char[0])
}
