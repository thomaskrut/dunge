package main

type point struct {
	x, y int
}

func (p point) getPossibleDirections(d *dungeonMap) map[direction]bool {
	directions := make(map[direction]bool)
	for _, dir := range getAllDirections() {
		newPoint := p
		newPoint.move(dir)
		if d.grid[newPoint.x][newPoint.y]&empty == empty {

			directions[dir] = true
		}
	}

	return directions
}

func getRandomPoint(d *dungeonMap) point {
	return point{x: randomNumber(len(d.grid)), y: randomNumber(len(d.grid[0]))}
}

func (p *point) move(dir direction) {

	p.x += dir.varX
	p.y += dir.varY

}

func (p point) isOutOfBounds(d *dungeonMap, margin int) bool {
	return p.x <= margin || p.x >= len(d.grid)-margin || p.y <= margin || p.y >= len(d.grid[0])-margin
}
