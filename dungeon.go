package main

import (
	"errors"
	"fmt"
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

func (d dungeon) print(charmap characterMapper, p player) {
	fmt.Println()
	for x, row := range d.grid {
		for y, cell := range row {
			if p.position.overlaps(Point{x: x, y: y}) {
				fmt.Printf("%c", p.getChar())
			} else {
				fmt.Printf("%c", charmap.chars[cell])
			}

		}
		println()
	}
	fmt.Println()
}

func (d *dungeon) setPoint(point Point, value int) {
	d.grid[point.x][point.y] = value
}

func (d *dungeon) getPoint(point Point) int {
	return d.grid[point.x][point.y]
}

func getEmptyPoint(d *dungeon) (point Point) {
	for {
		x := rand.Intn(len(d.grid))
		y := rand.Intn(len(d.grid[0]))
		if d.grid[x][y] == 1 {
			return Point{x, y}
		}
	}
}

func connectWithCorridor(d *dungeon, origin, destination Point) {
	currentPosition := origin
	var newDirection direction

	for {

		newDirection.directionTowards(currentPosition, destination).toNonDiagonal()

		currentPosition.move(newDirection)
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
	startingPoint.move(SouthEast)
	roomWidth := rand.Intn(maxWidth) + 4
	roomHeight := rand.Intn(maxHeight) + 4
	if p := (Point{x: startingPoint.x + roomWidth, y: startingPoint.y + roomHeight}); p.isOutOfBounds(d, 2) {
		return Point{}, errors.New("room out of bounds")
	}
	return d.createRoom(startingPoint, roomWidth, roomHeight)

}

func (d *dungeon) createRoom(startingPoint Point, width, height int) (position Point, err error) {

	for i := startingPoint.x; i < startingPoint.x+width; i++ {
		for j := startingPoint.y; j < startingPoint.y+height; j++ {
			if i == startingPoint.x+width-(width/2) {
				position.x = i
			}
			if j == startingPoint.y+height-(height/2) {
				position.y = j
			}

			if (d.getPoint(Point{x: i, y: j}) == empty) {
				return position, errors.New("space already empty")
			}
			d.setPoint(Point{x: i, y: j}, empty)
		}
	}
	return position, nil
}
