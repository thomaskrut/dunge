package main

import (
	"math/rand"
)

type Point struct {
	x, y int
}

func (p Point) getPossibleDirections(d *dungeon) map[direction]bool {
	directions := make(map[direction]bool)
	for _, dir := range getAllDirections() {
		newPoint := p
		newPoint.move(dir)
		if d.grid[newPoint.x][newPoint.y] & empty == empty {
			directions[dir] = true
		}
	}
	return directions
}

func (p1 Point) overlaps(p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y
}

func getRandomPoint(d *dungeon) Point {
	return Point{x: rand.Intn(len(d.grid)), y: rand.Intn(len(d.grid[0]))}
}

func (p *Point) move(dir direction) {
	p.x += dir.varX
	p.y += dir.varY
}

func (p Point) isOutOfBounds(d *dungeon, margin int) bool {
	return p.x <= margin || p.x >= len(d.grid)-margin || p.y <= margin || p.y >= len(d.grid[0])-margin
}
