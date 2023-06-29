package main

type feature struct {
	position point
	char string
	name string
	obstacle bool
}

func createDoor(position point) *feature {
	door := feature {
		position: position,
		char: "+",
		name: "door",
		obstacle: true,
	}
	d.grid[position.x][position.y] = wall
	return &door
}

func (f feature) getChar() rune {
	return rune(f.char[0])
}

