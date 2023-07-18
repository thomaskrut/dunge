package main

type point struct {
	X, Y int
}

func (p point) getPossibleDirections(d *dungeonMap) map[direction]bool {
	directions := make(map[direction]bool)
	for _, dir := range getAllDirections() {
		newPoint := p
		newPoint.move(dir)
		if d.read(newPoint)&empty == empty {

			directions[dir] = true
		}
	}

	return directions
}

func (p *point) move(dir direction) {

	p.X += dir.varX
	p.Y += dir.varY

}

func (p point) isOutOfBounds(margin int) bool {
	return p.X <= margin || p.X >= len(dungeon.grid)-margin || p.Y <= margin || p.Y >= len(dungeon.grid[0])-margin
}

func (p point) isInCorridor() bool {
	possibleDirections := p.getPossibleDirections(&dungeon)

	delete(possibleDirections, NorthEast)
	delete(possibleDirections, SouthEast)
	delete(possibleDirections, NorthEast)
	delete(possibleDirections, NorthWest)
	delete(possibleDirections, None)

	if len(possibleDirections) == 2 {

		if _, ok := possibleDirections[North]; ok {
			if _, ok := possibleDirections[South]; ok {
				return true
			}
		} else if _, ok := possibleDirections[East]; ok {
			if _, ok := possibleDirections[West]; ok {
				return true
			}
		}

	}
	return false

}
