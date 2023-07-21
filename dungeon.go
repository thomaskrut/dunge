package main

import (
	"errors"
)

type dungeon struct {
	Levels       map[int]*level
	Turn         int
	CurrentDepth int
}

type level struct {
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
		Levels:       make(map[int]*level),
		Turn:         0,
		CurrentDepth: 1,
	}
}

func (d *dungeon) newLevel(depth, width, height int) *level {
	zeroedGrid := make([][]byte, width)
	for i := range zeroedGrid {
		zeroedGrid[i] = make([]byte, height)
	}
	newLevel := level{
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

func (d *level) write(p point, value byte) {
	d.Grid[p.X][p.Y] = value
}

func (d *level) read(p point) byte {
	return d.Grid[p.X][p.Y]
}

func (d *level) getEmptyPoint() point {
	for {
		x := randomNumber(lev.Width)
		y := randomNumber(lev.Height)
		if lev.Grid[x][y] == empty {
			return point{x, y}
		}
	}
}

func (d *level) getPointInRoom() point {
	for {
		x := randomNumber(lev.Width)
		y := randomNumber(lev.Height)
		if lev.Grid[x][y]&room == room {
			return point{x, y}
		}
	}
}

func (d *level) getRandomPoint() point {
	return point{X: randomNumber(d.Width), Y: randomNumber(d.Height)}
}

func (d *level) generateItems(list itemList, numberOfIterations int) {

	for i := 1; i < numberOfIterations; i++ {

		rand := randomNumber(1000)

		for _, i := range list.Items {

			if rand < i.Prob {
				newItem := i
				newItem.setPosition(lev.getEmptyPoint())
				d.Items[newItem.Position] = append(d.Items[newItem.Position], &newItem)
			}
		}
	}

}

func (d *level) generateMonsters(list monsterList, numberOfIterations int) {

	for i := 0; i < numberOfIterations; i++ {

		rand := randomNumber(1000)

		for _, m := range list.Monsters {
			if rand < m.Prob {
				newMonster := m
				newMonster.setPosition(lev.getEmptyPoint())
				newMonster.Items = newInventory()
				newMonster.SpeedCounter = newMonster.Speed
				d.Monsters[newMonster.Position] = &newMonster
			}
		}
	}
}

func (l *level) excavate() {

	var previousRoom point
	var nextRoom point
	var err error

	for {
		previousRoom, err = l.newRoom(l.getRandomPoint(), 20, 20)
		if err != nil {
			continue
		}
		break
	}

	for i := 0; i < 10; i++ {

		for {
			nextRoom, err = l.newRoom(l.getRandomPoint(), 20, 20)
			if err != nil {
				continue
			}
			break
		}

		l.newCorridor(previousRoom, nextRoom)
		previousRoom = nextRoom

	}

}

func (l *level) newCorridor(origin, destination point) {
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

		if l.read(currentPosition)&room == room {
			l.write(currentPosition, empty|room)
		} else {
			l.write(currentPosition, empty)
		}

		if currentPosition == destination {
			break
		}
	}
}

func (l *level) newRoom(startingPoint point, maxWidth, maxHeight int) (position point, err error) {
	startingPoint.move(SouthEast)
	roomWidth := randomNumber(maxWidth) + 5
	roomHeight := randomNumber(maxHeight) + 5
	if p := (point{X: startingPoint.X + roomWidth, Y: startingPoint.Y + roomHeight}); p.isOutOfBounds(2) {
		return point{}, errors.New("room out of bounds")
	}
	return l.createRoom(startingPoint, roomWidth, roomHeight)

}

func (l *level) createRoom(startingPoint point, width, height int) (center point, err error) {

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
				if l.read(newPoint)&empty == empty {
					return center, errors.New("adjacent empty space")
				}
			}

			if i == startingPoint.X+width-1 {
				newPoint := currentPoint
				newPoint.move(East)
				if l.read(newPoint)&empty == empty {
					return center, errors.New("adjacent empty space")
				}
			}

			if j == startingPoint.Y {
				newPoint := currentPoint
				newPoint.move(North)
				if l.read(newPoint)&empty == empty {
					return center, errors.New("adjacent empty space")
				}
			}

			if j == startingPoint.Y+height-1 {
				newPoint := currentPoint
				newPoint.move(South)
				if l.read(newPoint)&empty == empty {
					return center, errors.New("adjacent empty space")
				}
			}

			if l.read(currentPoint) == empty {
				return center, errors.New("space already empty")
			}
			l.write(currentPoint, empty|room)
		}
	}
	return center, nil
}

func (l *level) generateDoors(numberOfDoors int) {

	for i := 0; i < numberOfDoors; {

		p := lev.getEmptyPoint()

		if door, ok := createDoor(p); ok {
			if door.State == "closed" {
				lev.write(door.Position, obstacle)
			}
			lev.Features[p] = door
			i++
		}

	}

}

func (l *level) generateStairs() {

	if world.CurrentDepth > 1 {
		stairs, _ := createStairs(pl.Position, "up")
		lev.Features[pl.Position] = stairs
		lev.Upstair = pl.Position
	}

	for {
		newPoint := lev.getEmptyPoint()
		if stairs, ok := createStairs(newPoint, "down"); ok {
			lev.Features[newPoint] = stairs
			lev.Downstair = newPoint
			break
		}
	}

}
