package main

type feature struct {
	Position    point
	Prefix      string
	Name        string
	Char        string
	Description string
	Closed      bool
}

func createDoor(position point) (*feature, bool) {

	if position.isInCorridor() {
		door := feature{
			Position: position,
			Prefix:   "a",
			Name:     "door",
		}
		if randomNumber(2) == 1 {
			door.Closed = true
			door.Char = "+"
			door.Description = "a closed door"
		} else {
			door.Closed = false
			door.Char = "-"
			door.Description = "an open door"
		}
		return &door, true
	} else {
		return nil, false
	}

}

func createStairs(position point, direction string) (*feature, bool) {
	if _, ok := featuresOnMap[position]; ok || position.isInCorridor() {
		return nil, false
	}

	stairs := feature {
		Position: position,
	}

	switch direction {
	case "up":
		stairs.Char = "<"
		stairs.Name = "upstair"
		stairs.Prefix = "a"
		stairs.Description = "a staircase going up"
	case "down":
		stairs.Char = ">"
		stairs.Name = "downstair"
		stairs.Prefix = "a"
		stairs.Description = "a staircase going down"
	}
	
	return &stairs, true

}

func (f feature) getChar() rune {
	return rune(f.Char[0])
}
