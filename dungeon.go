package main

import (
	"errors"
)

type dungeonMap struct {
	grid          [][]int
	width, height int
}

func newDungeon(width, height int) dungeonMap {
	zeroedGrid := make([][]int, width)
	for i := range zeroedGrid {
		zeroedGrid[i] = make([]int, height)
	}
	return dungeonMap{grid: zeroedGrid, width: width, height: height}
}

func (d *dungeonMap) write(p point, value int) {
	d.grid[p.x][p.y] = value
}

func (d *dungeonMap) read(p point) int {
	return d.grid[p.x][p.y]
}

func getEmptyPoint(d *dungeonMap) point {
	for {
		x := randomNumber(len(d.grid))
		y := randomNumber(len(d.grid[0]))
		if d.grid[x][y] == empty {
			return point{x, y}
		}
	}
}

func connectWithCorridor(d *dungeonMap, origin, destination point) {
	currentPosition := origin
	var newDirection direction

	for {

		if randomNumber(2) == 0 {
			newDirection.connect(currentPosition, destination).toNonDiagonal()
		}

		currentPosition.move(newDirection)
		if currentPosition.isOutOfBounds(d, 2) {
			break
		}
		if d.grid[currentPosition.x][currentPosition.y]&room == room {
			d.write(currentPosition, empty|room)
		} else {
			d.write(currentPosition, empty)
		}

		if currentPosition == destination {
			break
		}
	}
}

func (d *dungeonMap) createRandomRoom(startingPoint point, maxWidth, maxHeight int) (position point, err error) {
	startingPoint.move(SouthEast)
	roomWidth := randomNumber(maxWidth) + 3
	roomHeight := randomNumber(maxHeight) + 3
	if p := (point{x: startingPoint.x + roomWidth, y: startingPoint.y + roomHeight}); p.isOutOfBounds(d, 2) {
		return point{}, errors.New("room out of bounds")
	}
	return d.createRoom(startingPoint, roomWidth, roomHeight)

}

func (d *dungeonMap) createRoom(startingPoint point, width, height int) (center point, err error) {

	for i := startingPoint.x; i < startingPoint.x+width; i++ {
		for j := startingPoint.y; j < startingPoint.y+height; j++ {
			if i == startingPoint.x+width-(width/2) {
				center.x = i
			}
			if j == startingPoint.y+height-(height/2) {
				center.y = j
			}

			if (d.read(point{x: i, y: j}) == empty) {
				return center, errors.New("space already empty")
			}
			d.write(point{x: i, y: j}, empty|room)
		}
	}
	return center, nil
}

func (d *dungeonMap) generateDoors(numberOfDoors int) {

	count := 0

	for count < numberOfDoors {

		p := getEmptyPoint(d)

		possibleDirections := p.getPossibleDirections(d)

		delete(possibleDirections, NorthEast)
		delete(possibleDirections, SouthEast)
		delete(possibleDirections, NorthEast)
		delete(possibleDirections, NorthWest)
		delete(possibleDirections, None)

		if len(possibleDirections) == 2 {

			if _, ok := possibleDirections[North]; ok {
				if _, ok := possibleDirections[South]; ok {
					featuresOnMap[p] = createDoor(p)
					count++
				}
			} else if _, ok := possibleDirections[East]; ok {
				if _, ok := possibleDirections[West]; ok {
					featuresOnMap[p] = createDoor(p)
					count++
				}
			}

		}

	}

}
