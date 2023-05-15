package main

type point struct {
	x, y int
}

func (p point) getPossibleDirections(d *dungeon) map[direction]bool {
	directions := make(map[direction]bool)
	for _, dir := range getAllDirections() {
		newPoint := p
		newPoint.new(dir)
		if d.grid[newPoint.x][newPoint.y]&empty == empty {
			directions[dir] = true
		}
	}

	return directions
}

func (p1 point) overlaps(p2 point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func getRandomPoint(d *dungeon) point {
	return point{x: randomNumber(len(d.grid)), y: randomNumber(len(d.grid[0]))}
}

func (p *point) new(dir direction) {

	p.x += dir.varX
	p.y += dir.varY

}

func (p point) isOutOfBounds(d *dungeon, margin int) bool {
	return p.x <= margin || p.x >= len(d.grid)-margin || p.y <= margin || p.y >= len(d.grid[0])-margin
}
