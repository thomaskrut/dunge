package main

import (
	"errors"
)

type dungeon struct {
	Levels map[int]*levelMap
	Turn         int
	CurrentDepth int
}

type levelMap struct {
	Grid          [][]byte
	Upstair       point
	Downstair     point
	Width, Height int
	Monsters      map[point]*monster
	Items         map[point][]*item
	Features      map[point]*feature
}

func newDungeon() dungeon {
	return dungeon{
		Levels:       make(map[int]*levelMap),
		Turn:         0,
		CurrentDepth: 1,
	}
}

func (d *dungeon) newLevel(depth, width, height int) *levelMap {
	zeroedGrid := make([][]byte, width)
	for i := range zeroedGrid {
		zeroedGrid[i] = make([]byte, height)
	}
	newLevel := levelMap{
		Grid:     zeroedGrid,
		Width:    width,
		Height:   height,
		Monsters: make(map[point]*monster),
		Items:    make(map[point][]*item),
		Features: make(map[point]*feature),
	}
	d.Levels[depth] = &newLevel
	return &newLevel
}

func (d *levelMap) write(p point, value byte) {
	d.Grid[p.X][p.Y] = value
}

func (d *levelMap) read(p point) byte {
	return d.Grid[p.X][p.Y]
}

func (d *levelMap) getEmptyPoint() point {
	for {
		x := randomNumber(level.Width)
		y := randomNumber(level.Height)
		if level.Grid[x][y] == empty {
			return point{x, y}
		}
	}
}

func (d *levelMap) getPointInRoom() point {
	for {
		x := randomNumber(level.Width)
		y := randomNumber(level.Height)
		if level.Grid[x][y]&room == room {
			return point{x, y}
		}
	}
}

func (d *levelMap) getRandomPoint() point {
	return point{X: randomNumber(d.Width), Y: randomNumber(d.Height)}
}

func (d *levelMap) generateItems(list itemList, numberOfIterations int) {

	for i := 1; i < numberOfIterations; i++ {

		rand := randomNumber(1000)

		for _, i := range list.Items {

			if rand < i.Prob {
				newItem := i
				newItem.setPosition(level.getEmptyPoint())
				d.Items[newItem.Position] = append(d.Items[newItem.Position], &newItem)
			}
		}
	}

}

func (d *levelMap) generateMonsters(list monsterList, numberOfIterations int) {

	for i := 0; i < numberOfIterations; i++ {

		rand := randomNumber(1000)

		for _, m := range list.Monsters {
			if rand < m.Prob {
				newMonster := m
				newMonster.setPosition(level.getEmptyPoint())
				newMonster.Items = newInventory()
				newMonster.SpeedCounter = newMonster.Speed
				d.Monsters[newMonster.Position] = &newMonster
			}
		}
	}
}

func (d *levelMap) newCorridor(origin, destination point) {
	currentPosition := origin
	var newDirection direction

	for {

		if randomNumber(2) == 0 {
			newDirection.connect(currentPosition, destination).toNonDiagonal()
		}
		
		currentPosition.move(newDirection)
		if currentPosition.isOutOfBounds(2) {
			break
		}

		if d.read(currentPosition)&room == room {
			d.write(currentPosition, empty|room)
		} else {
			d.write(currentPosition, empty)
		}

		if currentPosition == destination {
			break
		}
	}
}

func (d *levelMap) newRoom(startingPoint point, maxWidth, maxHeight int) (position point, err error) {
	startingPoint.move(SouthEast)
	roomWidth := randomNumber(maxWidth) + 5
	roomHeight := randomNumber(maxHeight) + 5
	if p := (point{X: startingPoint.X + roomWidth, Y: startingPoint.Y + roomHeight}); p.isOutOfBounds(2) {
		return point{}, errors.New("room out of bounds")
	}
	return d.createRoom(startingPoint, roomWidth, roomHeight)

}

func (d *levelMap) createRoom(startingPoint point, width, height int) (center point, err error) {

	for i := startingPoint.X; i < startingPoint.X+width; i++ {
		for j := startingPoint.Y; j < startingPoint.Y+height; j++ {
			currentPoint := point{X: i, Y: j}
			if i == startingPoint.X+width-(width/2) {
				center.X = i
			}
			if j == startingPoint.Y+height-(height/2) {
				center.Y = j
			}

			if i == startingPoint.X {
				newPoint := currentPoint
				newPoint.move(West)
				if d.read(newPoint)&empty == empty {
					return center, errors.New("adjacent empty space")
				}
			}

			if i == startingPoint.X+width - 1 {
				newPoint := currentPoint
				newPoint.move(East)
				if d.read(newPoint)&empty == empty {
					return center, errors.New("adjacent empty space")
				}
			}

			if j == startingPoint.Y {
				newPoint := currentPoint
				newPoint.move(North)
				if d.read(newPoint)&empty == empty {
					return center, errors.New("adjacent empty space")
				}
			}

			if j == startingPoint.Y+height-1 {
				newPoint := currentPoint
				newPoint.move(South)
				if d.read(newPoint)&empty == empty {
					return center, errors.New("adjacent empty space")
				}
			}

			if d.read(currentPoint) == empty {
				return center, errors.New("space already empty")
			}
			d.write(currentPoint, empty|room)
		}
	}
	return center, nil
}

func (d *levelMap) generateDoors(numberOfDoors int) {

	for i := 0; i < numberOfDoors; {

		p := level.getEmptyPoint()

		if door, ok := createDoor(p); ok {
			if door.State == "closed" {
				level.write(door.Position, obstacle)
			}
			level.Features[p] = door
			i++
		}

	}

}

func (d *levelMap) generateStairs() {

	if world.CurrentDepth > 1 {
		stairs, _ := createStairs(p.Position, "up")
		level.Features[p.Position] = stairs
		level.Upstair = p.Position
	}

	for {
		newPoint := level.getEmptyPoint()
		if stairs, ok := createStairs(newPoint, "down"); ok {
			level.Features[newPoint] = stairs
			level.Downstair = newPoint
			break
		}
	}

}
