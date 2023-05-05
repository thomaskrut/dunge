package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Dungeon struct {
	grid [][]int
}

func newDungeon(width, height int) Dungeon {
	zeroedGrid := make([][]int, height)
	for i := range zeroedGrid {
		zeroedGrid[i] = make([]int, width)
	}
	return Dungeon{grid: zeroedGrid}
}

func (d Dungeon) print(charmap characterMapper) {
	fmt.Println()
	for _, row := range d.grid {
		for _, cell := range row {
			fmt.Printf("%c", charmap.chars[cell])
		}
		println()
	}
	fmt.Println()
}

func (d *Dungeon) setPoint(point Point, value int) {
	d.grid[point.x][point.y] = value
}

func (d *Dungeon) getPoint(point Point) int {
	return d.grid[point.x][point.y]
}

func getEmptyPoint(d *Dungeon) (point Point) {
	for {
		x := rand.Intn(len(d.grid))
		y := rand.Intn(len(d.grid[0]))
		if d.grid[x][y] == 1 {
			return Point{x, y}
		}
	}
}

func connectWithCorridor(d *Dungeon, origin, destination Point) {
	currentPosition := origin;
	var newDirection Direction

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

func (d *Dungeon) createCorridor(maxLength int) (endPoint Point) {

	currentPosition := getEmptyPoint(d)
	dir := randomDirection(None, false, false)
	length := rand.Intn(maxLength) + 10
	for length > 0 {

		if rand.Intn(6) == 1 {
			dir = randomDirection(dir, false, false)
		}
		
		testPosition := currentPosition
		testPosition.move(dir)

		for testPosition.isOutOfBounds(d, 2) {
			
			dir = randomDirection(dir, false, false)
			testPosition = currentPosition
			testPosition.move(dir)
			
		}
		currentPosition.move(dir)
		d.setPoint(currentPosition, empty)
		length--
	}
	return currentPosition

}

func (d *Dungeon) createRandomRoom(startingPoint Point, maxWidth, maxHeight int) (position Point, err error) {
		startingPoint.move(SouthEast)
		roomWidth := rand.Intn(maxWidth) + 4
		roomHeight := rand.Intn(maxHeight) + 4
		if p := (Point{x: startingPoint.x + roomWidth, y: startingPoint.y + roomHeight}); p.isOutOfBounds(d, 2) {
			return Point{}, errors.New("room out of bounds")
		}
		return d.createRoom(startingPoint, roomWidth, roomHeight)
		
}

func (d *Dungeon) createRoom(startingPoint Point, width, height int) (position Point, err error) {
	
	for i := startingPoint.x; i < startingPoint.x+width; i++ {
		for j := startingPoint.y; j < startingPoint.y+height; j++ {
			if i == startingPoint.x + width - (width / 2) {
				position.x = i
			}
			if j == startingPoint.y + height - (height / 2) {
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
