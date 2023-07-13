package main

import (
	"errors"
)

type dungeonMap struct {
	grid          [][]byte
	width, height int
}

func newDungeon(width, height int) dungeonMap {
	zeroedGrid := make([][]byte, width)
	for i := range zeroedGrid {
		zeroedGrid[i] = make([]byte, height)
	}
	return dungeonMap{grid: zeroedGrid, width: width, height: height}
}

func (d *dungeonMap) write(p point, value byte) {
	d.grid[p.x][p.y] = value
}

func (d *dungeonMap) read(p point) byte {
	return d.grid[p.x][p.y]
}

func (d *dungeonMap) getEmptyPoint() point {
	for {
		x := randomNumber(dungeon.width)
		y := randomNumber(dungeon.height)
		if dungeon.grid[x][y] == empty {
			return point{x, y}
		}
	}
}

func (d *dungeonMap) getRandomPoint() point {
	return point{x: randomNumber(d.width), y: randomNumber(d.height)}
}


func (d *dungeonMap) connectWithCorridor(origin, destination point) {
	currentPosition := origin
	//var previousPosition point
	var newDirection direction

	for {

		if randomNumber(2) == 0 {
			newDirection.connect(currentPosition, destination).toNonDiagonal()
		}
		//previousPosition = currentPosition
		currentPosition.move(newDirection)
		if currentPosition.isOutOfBounds(2) {
			break
		}

		if d.read(currentPosition)&room == room {
			d.write(currentPosition, empty|room)
		} else {
			d.write(currentPosition, empty)
		}

		/*if d.read(previousPosition)&room != room && d.read(currentPosition)&room == room {
			if door, ok := createDoor(previousPosition); ok {
				featuresOnMap[previousPosition] = door
			}
		}

		if d.read(previousPosition)&room == room && d.read(currentPosition)&room != room {
			if door, ok := createDoor(currentPosition); ok {
				featuresOnMap[currentPosition] = door
			}
		}*/

		if currentPosition == destination {
			break
		}
	}
}

func (d *dungeonMap) createRandomRoom(startingPoint point, maxWidth, maxHeight int) (position point, err error) {
	startingPoint.move(SouthEast)
	roomWidth := randomNumber(maxWidth) + 5
	roomHeight := randomNumber(maxHeight) + 5
	if p := (point{x: startingPoint.x + roomWidth, y: startingPoint.y + roomHeight}); p.isOutOfBounds(2) {
		return point{}, errors.New("room out of bounds")
	}
	return d.createRoom(startingPoint, roomWidth, roomHeight)

}

func (d *dungeonMap) createRoom(startingPoint point, width, height int) (center point, err error) {

	for i := startingPoint.x; i < startingPoint.x+width; i++ {
		for j := startingPoint.y; j < startingPoint.y+height; j++ {
			currentPoint := point{x: i, y: j}
			if i == startingPoint.x+width-(width/2) {
				center.x = i
			}
			if j == startingPoint.y+height-(height/2) {
				center.y = j
			}

			if d.read(currentPoint) == empty {
				return center, errors.New("space already empty")
			}
			d.write(currentPoint, empty|room)
		}
	}
	return center, nil
}

func (d *dungeonMap) generateDoors(numberOfDoors int) {

	count := 0

	for count < numberOfDoors {

		p := dungeon.getEmptyPoint()

		if door, ok := createDoor(p); ok {
			if door.closed {
				dungeon.write(door.position, obstacle)
			}
			featuresOnMap[p] = door
			count++
		}

	}

}
