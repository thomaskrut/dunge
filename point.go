package main

type point struct {
	X, Y int
}

func (p point) getPossibleDirections(d *level) map[direction]bool {
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
	return p.X <= margin || p.X >= len(lev.Grid)-margin || p.Y <= margin || p.Y >= len(lev.Grid[0])-margin
}

func (p point) isInCorridor() bool {
	possibleDirections := p.getPossibleDirections(lev)

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
