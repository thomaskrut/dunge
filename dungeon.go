package main

import (
	"errors"
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

func (d *dungeon) setPoint(p point, value int) {
	d.grid[p.x][p.y] = value
}

func (d *dungeon) getPoint(p point) int {
	return d.grid[p.x][p.y]
}

func getEmptyPoint(d *dungeon) point {
	for {
		x := randomNumber(len(d.grid))
		y := randomNumber(len(d.grid[0]))
		if d.grid[x][y] == empty {
			return point{x, y}
		}
	}
}

func connectWithCorridor(d *dungeon, origin, destination point) {
	currentPosition := origin
	var newDirection direction

	for {

		if randomNumber(2) == 0 {
			newDirection.connect(currentPosition, destination).toNonDiagonal()
		}

		currentPosition.move(newDirection)
		if currentPosition.isOutOfBounds(d, 1) {
			break
		}
		d.setPoint(currentPosition, empty)
		if currentPosition == destination {
			break
		}
	}
}

func (d *dungeon) createRandomRoom(startingPoint point, maxWidth, maxHeight int) (position point, err error) {
	startingPoint.move(SouthEast)
	roomWidth := randomNumber(maxWidth) + 3
	roomHeight := randomNumber(maxHeight) + 3
	if p := (point{x: startingPoint.x + roomWidth, y: startingPoint.y + roomHeight}); p.isOutOfBounds(d, 2) {
		return point{}, errors.New("room out of bounds")
	}
	return d.createRoom(startingPoint, roomWidth, roomHeight)

}

func (d *dungeon) createRoom(startingPoint point, width, height int) (center point, err error) {

	for i := startingPoint.x; i < startingPoint.x+width; i++ {
		for j := startingPoint.y; j < startingPoint.y+height; j++ {
			if i == startingPoint.x+width-(width/2) {
				center.x = i
			}
			if j == startingPoint.y+height-(height/2) {
				center.y = j
			}

			if (d.getPoint(point{x: i, y: j}) == empty) {
				return center, errors.New("space already empty")
			}
			d.setPoint(point{x: i, y: j}, empty)
		}
	}
	return center, nil
}
