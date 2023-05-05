package main

import (
	"math/rand"
)

type Point struct {
	x, y int
}

func (p1 Point) overlaps(p2 Point) bool {
	return p1.x == p2.x && p1.y == p2.y 
}

func getRandomPoint(d *Dungeon) Point {
	return Point{x: rand.Intn(len(d.grid)), y: rand.Intn(len(d.grid[0]))}
}

func (p *Point) move(dir Direction) {
	p.x += dir.varX
	p.y += dir.varY
}

func (p Point) isOutOfBounds(d *Dungeon, margin int) bool {
	return p.x <= margin || p.x >= len(d.grid) - margin || p.y <= margin || p.y >= len(d.grid[0]) - margin
}
