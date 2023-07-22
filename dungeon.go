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

func (l *level) write(p point, value byte) {
	l.Grid[p.X][p.Y] = value
}

func (l *level) read(p point) byte {
	return l.Grid[p.X][p.Y]
}

func (l *level) getEmptyPoint() point {
	for {
		x := randomNumber(lvl.Width)
		y := randomNumber(lvl.Height)
		if lvl.Grid[x][y] == empty {
			return point{x, y}
		}
	}
}

func (l *level) getPointInRoom() point {
	for {
		x := randomNumber(lvl.Width)
		y := randomNumber(lvl.Height)
		if lvl.Grid[x][y]&room == room {
			return point{x, y}
		}
	}
}

func (l *level) getRandomPoint() point {
	return point{X: randomNumber(l.Width), Y: randomNumber(l.Height)}
}

func (l *level) generateItems(list itemList, numberOfIterations int) {

	for i := 1; i < numberOfIterations; i++ {

		rand := randomNumber(1000)

		for _, i := range list.Items {

			if rand < i.Prob {
				newItem := i
				newItem.setPosition(lvl.getEmptyPoint())
				l.Items[newItem.Position] = append(l.Items[newItem.Position], &newItem)
			}
		}
	}

}

func (l *level) moveMonsters() {

	for i := 0; i < pl.Speed; i++ {

		for i, m := range l.Monsters {

			if m.readyToMove() {

				if items, ok := l.Items[m.Position]; ok && m.CarriesItems && randomNumber(20) > m.Speed {
					if l.read(m.Position)&lit == lit {
						messages.push("The "+m.Name+" picked up "+items[len(items)-1].Prefix+" "+items[len(items)-1].Name, gameplay)
					}
					m.Items.add(items[len(items)-1])
					l.Items[m.Position] = l.Items[m.Position][:len(l.Items[m.Position])-1]
					if len(l.Items[m.Position]) == 0 {
						delete(l.Items, m.Position)
					}
					continue
				}

				var newDirection direction
				newDirection.connect(m.getPosition(), pl.getPosition())

				if !m.Aggressive {
					newDirection = newDirection.opposite()
				}

				if !m.MovesDiagonally {
					newDirection.toNonDiagonal()
				}
				for i := 0; !m.attemptMove(newDirection) && i < 10; i++ {
					newDirection = randomDirection(newDirection, false, m.MovesDiagonally)
				}
				delete(l.Monsters, i)
				l.Monsters[m.Position] = m
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
				newMonster.setPosition(lvl.getEmptyPoint())
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
		previousRoom, err = l.newRoom(l.getRandomPoint(), 14, 14)
		if err != nil {
			continue
		}
		break
	}

	for i := 0; i < (width+height)/20; i++ {

		for {
			nextRoom, err = l.newRoom(l.getRandomPoint(), 14, 14)
			if err != nil {
				continue
			}
			break
		}

		l.newCorridor(previousRoom, nextRoom)

		previousRoom = nextRoom

	}

	l.newCorridor(l.getEmptyPoint(), l.getEmptyPoint())
	l.newCorridor(l.getEmptyPoint(), l.getEmptyPoint())
	

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
	newRoom, centerPoint, err := l.createRoom(startingPoint, roomWidth, roomHeight)
	if err != nil {
		return point{}, errors.New("room out of bounds")
	}
	for _, p := range newRoom {
		l.write(p, empty|room)
	}
	return centerPoint, nil

}

func (l *level) createRoom(startingPoint point, width, height int) (newRoom []point, centerPoint point, err error) {

	for i := startingPoint.X; i < startingPoint.X+width; i++ {
		for j := startingPoint.Y; j < startingPoint.Y+height; j++ {
			currentPoint := point{X: i, Y: j}
			if i == startingPoint.X+width-(width/2) {
				centerPoint.X = i
			}
			if j == startingPoint.Y+height-(height/2) {
				centerPoint.Y = j
			}

			if i == startingPoint.X {
				newPoint := currentPoint
				newPoint.move(West)
				if l.read(newPoint) == empty || l.read(newPoint)&room == room {
					return nil, centerPoint, errors.New("adjacent empty space")
				}
			}

			if i == startingPoint.X+width-1 {
				newPoint := currentPoint
				newPoint.move(East)
				if l.read(newPoint) == empty || l.read(newPoint)&room == room {
					return nil, centerPoint, errors.New("adjacent empty space")
				}
			}

			if j == startingPoint.Y {
				newPoint := currentPoint
				newPoint.move(North)
				if l.read(newPoint) == empty || l.read(newPoint)&room == room {
					return nil, centerPoint, errors.New("adjacent empty space")
				}
			}

			if j == startingPoint.Y+height-1 {
				newPoint := currentPoint
				newPoint.move(South)
				if l.read(newPoint) == empty || l.read(newPoint)&room == room {
					return nil, centerPoint, errors.New("adjacent empty space")
				}
			}

			if l.read(currentPoint) == empty {
				return nil, centerPoint, errors.New("space already empty")
			}
			newRoom = append(newRoom, currentPoint)
		}
	}
	return newRoom, centerPoint, nil
}

func (l *level) generateDoors(numberOfDoors int) {

	for i := 0; i < numberOfDoors; {

		p := lvl.getEmptyPoint()

		if door, ok := createDoor(p); ok {
			if door.State == "closed" {
				lvl.write(door.Position, obstacle)
			}
			lvl.Features[p] = door
			i++
		}

	}

}

func (l *level) generateStairs() {

	if world.CurrentDepth > 1 {
		stairs, _ := createStairs(pl.Position, "up")
		lvl.Features[pl.Position] = stairs
		lvl.Upstair = pl.Position
	}

	for {
		newPoint := lvl.getEmptyPoint()
		if stairs, ok := createStairs(newPoint, "down"); ok {
			lvl.Features[newPoint] = stairs
			lvl.Downstair = newPoint
			break
		}
	}

}
