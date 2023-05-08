package main

type player struct {
	position    Point
	char        rune
	lightsource int
}

func (p *player) move(dir direction) {
	if p.position.getPossibleDirections(&d)[dir] {
		alterArea(&d, p.position, visited, p.lightsource)
		p.position.move(dir)
		alterArea(&d, p.position, lit, p.lightsource)
	}
}

func alterArea(d *dungeon, p Point, value int, currentDepth int) {
	if currentDepth == 0 {
		return
	}
	for _, dir := range getNonDiagonalDirections() {
		newPoint := p
		newPoint.move(dir)
		if d.grid[newPoint.x][newPoint.y] >= empty {
			d.grid[newPoint.x][newPoint.y] = value
			alterArea(d, newPoint, value, currentDepth-1)
		} else {
			d.grid[newPoint.x][newPoint.y] = 0
		}
	}

}

func (p player) getPosition() Point {
	return p.position
}

func (p *player) setPosition(point Point) {
	p.position = point
}

func (p player) getChar() rune {
	return p.char
}

func newPlayer() player {
	return player{char: '@', lightsource: 5}
}
