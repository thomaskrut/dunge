package main

import (
	"errors"
)

type dungeonMap struct {
	Grid          [][]byte
	Width, Height int
}

func newDungeon(width, height int) dungeonMap {
	zeroedGrid := make([][]byte, width)
	for i := range zeroedGrid {
		zeroedGrid[i] = make([]byte, height)
	}
	return dungeonMap{Grid: zeroedGrid, Width: width, Height: height}
}

func (d *dungeonMap) write(p point, value byte) {
	d.Grid[p.X][p.Y] = value
}

func (d *dungeonMap) read(p point) byte {
	return d.Grid[p.X][p.Y]
}

func (d *dungeonMap) getEmptyPoint() point {
	for {
		x := randomNumber(dungeon.Width)
		y := randomNumber(dungeon.Height)
		if dungeon.Grid[x][y] == empty {
			return point{x, y}
		}
	}
}

func (d *dungeonMap) getPointInRoom() point {
	for {
		x := randomNumber(dungeon.Width)
		y := randomNumber(dungeon.Height)
		if dungeon.Grid[x][y]&room == room {
			return point{x, y}
		}
	}
}

func (d *dungeonMap) getRandomPoint() point {
	return point{X: randomNumber(d.Width), Y: randomNumber(d.Height)}
}

func (d *dungeonMap) newCorridor(origin, destination point) {
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

func (d *dungeonMap) newRoom(startingPoint point, maxWidth, maxHeight int) (position point, err error) {
	startingPoint.move(SouthEast)
	roomWidth := randomNumber(maxWidth) + 5
	roomHeight := randomNumber(maxHeight) + 5
	if p := (point{X: startingPoint.X + roomWidth, Y: startingPoint.Y + roomHeight}); p.isOutOfBounds(2) {
		return point{}, errors.New("room out of bounds")
	}
	return d.createRoom(startingPoint, roomWidth, roomHeight)

}

func (d *dungeonMap) createRoom(startingPoint point, width, height int) (center point, err error) {

	for i := startingPoint.X; i < startingPoint.X+width; i++ {
		for j := startingPoint.Y; j < startingPoint.Y+height; j++ {
			currentPoint := point{X: i, Y: j}
			if i == startingPoint.X+width-(width/2) {
				center.X = i
			}
			if j == startingPoint.Y+height-(height/2) {
				center.Y = j
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

	for i := 0; i < numberOfDoors; {

		p := dungeon.getEmptyPoint()

		if door, ok := createDoor(p); ok {
			if door.State == "closed" {
				dungeon.write(door.Position, obstacle)
			}
			featuresOnMap[p] = door
			i++
		}

	}

}

func (d *dungeonMap) generateStairs(up, down int) {

	for i := 0; i < up; {
		p := dungeon.getEmptyPoint()
		if stairs, ok := createStairs(p, "up"); ok {
			featuresOnMap[p] = stairs
			i++
		}
	}

	for i := 0; i < down; {
		p := dungeon.getEmptyPoint()
		if stairs, ok := createStairs(p, "down"); ok {
			featuresOnMap[p] = stairs
			i++
		}
	}

}
