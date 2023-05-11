package main

import (
	"errors"
	"math/rand"
)

type dungeon struct {
	grid          [][]int
	width, height int
}

func newDungeon(width, height int) dungeon {
	zeroedGrid := make([][]int, width)
	for i := range zeroedGrid {
		zeroedGrid[i] = make([]int, height)
	}
	return dungeon{grid: zeroedGrid, width: width, height: height}
}

func (d *dungeon) setPoint(p Point, value int) {
	d.grid[p.x][p.y] = value
}

func (d *dungeon) getPoint(p Point) int {
	return d.grid[p.x][p.y]
}

func getEmptyPoint(d *dungeon) Point {
	for {
		x := rand.Intn(len(d.grid))
		y := rand.Intn(len(d.grid[0]))
		if d.grid[x][y] == empty {
			return Point{x, y}
		}
	}
}

func connectWithCorridor(d *dungeon, origin, destination Point) {
	currentPosition := origin
	var newDirection direction

	for {

		newDirection.connect(currentPosition, destination).toNonDiagonal()

		currentPosition.new(newDirection)
		if currentPosition.isOutOfBounds(d, 1) {
			break
		}
		d.setPoint(currentPosition, empty)
		if currentPosition.overlaps(destination) {
			break
		}
	}
}

func (d *dungeon) createRandomRoom(startingPoint Point, maxWidth, maxHeight int) (position Point, err error) {
	startingPoint.new(SouthEast)
	roomWidth := rand.Intn(maxWidth) + 4
	roomHeight := rand.Intn(maxHeight) + 4
	if p := (Point{x: startingPoint.x + roomWidth, y: startingPoint.y + roomHeight}); p.isOutOfBounds(d, 2) {
		return Point{}, errors.New("room out of bounds")
	}
	return d.createRoom(startingPoint, roomWidth, roomHeight)

}

func (d *dungeon) createRoom(startingPoint Point, width, height int) (center Point, err error) {

	for i := startingPoint.x; i < startingPoint.x+width; i++ {
		for j := startingPoint.y; j < startingPoint.y+height; j++ {
			if i == startingPoint.x+width-(width/2) {
				center.x = i
			}
			if j == startingPoint.y+height-(height/2) {
				center.y = j
			}

			if (d.getPoint(Point{x: i, y: j}) == empty) {
				return center, errors.New("space already empty")
			}
			d.setPoint(Point{x: i, y: j}, empty)
		}
	}
	return center, nil
}
