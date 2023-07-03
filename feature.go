package main

type feature struct {
	position point
	char     string
	name     string
	closed   bool
}

func createDoor(position point) (*feature, bool) {

	possibleDirections := position.getPossibleDirections(&dungeon)

	delete(possibleDirections, NorthEast)
	delete(possibleDirections, SouthEast)
	delete(possibleDirections, NorthEast)
	delete(possibleDirections, NorthWest)
	delete(possibleDirections, None)

	if len(possibleDirections) == 2 {

		door := feature{
			position: position,
			name:     "door",
		}
		if randomNumber(2) == 1 {
			door.closed = true
			door.char = "+"
		} else {
			door.closed = false
			door.char = "-"
		}

		if _, ok := possibleDirections[North]; ok {
			if _, ok := possibleDirections[South]; ok {
				return &door, true
			}
		} else if _, ok := possibleDirections[East]; ok {
			if _, ok := possibleDirections[West]; ok {
				return &door, true
			}
		}

	}
	return &feature{}, false
}

func (f feature) getChar() rune {
	return rune(f.char[0])
}
