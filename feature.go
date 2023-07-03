package main

type feature struct {
	position point
	char     string
	name     string
	obstacle bool
}

func createDoor(position point) *feature {
	door := feature{
		position: position,
		name:     "door",
	}
	if randomNumber(2) == 1 {
		door.obstacle = true
		door.char = "+"
		dungeon.grid[position.x][position.y] = obstacle
	} else {
		door.obstacle = false
		door.char = "-"
	}

	return &door
}

func (f feature) getChar() rune {
	return rune(f.char[0])
}
