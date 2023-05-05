package main

import (
	"math/rand"
)

var (
	None      = Direction{varX: 0, varY: 0}
	North     = Direction{varX: 0, varY: -1}
	South     = Direction{varX: 0, varY: 1}
	East      = Direction{varX: 1, varY: 0}
	West      = Direction{varX: -1, varY: 0}
	NorthWest = Direction{varX: -1, varY: -1}
	NorthEast = Direction{varX: 1, varY: -1}
	SouthWest = Direction{varX: -1, varY: 1}
	SouthEast = Direction{varX: 1, varY: 1}
)

type Direction struct {
	varX, varY int
}

func (d *Direction) toNonDiagonal() {
	r := rand.Intn(2)
	switch *d {
	case NorthWest:
		switch r {
		case 0:
			*d = North
		case 1:
			*d = West
		}
	case NorthEast:
		switch r {
		case 0:
			*d = North
		case 1:
			*d = East
		}
	case SouthWest:
		switch r {
		case 0:
			*d = South
		case 1:
			*d = West
		}
	case SouthEast:
		switch r {
		case 0:
			*d = East
		case 1:
			*d = South
		}

	}
	
}

func (dir *Direction) directionTowards(from, to Point) (newDir *Direction) {
	switch {
	case to.x == from.x && to.y < from.y:
		*dir = North
	case to.x == from.x && to.y > from.y:
		*dir = South
	case to.x < from.x && to.y == from.y:
		*dir = West
	case to.x > from.x && to.y == from.y:
		*dir = East
	case to.x > from.x && to.y < from.y:
		*dir = NorthEast
	case to.x > from.x && to.y > from.y:
		*dir = SouthEast
	case to.x < from.x && to.y < from.y:
		*dir = NorthWest
	case to.x < from.x && to.y > from.y:
		*dir = SouthWest
	}
	return dir
	
}

func randomDirection(currentDirection Direction, allowOpposite bool, allowDiagonal bool) Direction {
	var r int
	var newDirection Direction
	if allowDiagonal {
		r = 8
	} else {
		r = 4
	}
	switch rand.Intn(r) {
	case 0:
		newDirection = North
	case 1:
		newDirection = South
	case 2:
		newDirection = West
	case 3:
		newDirection = East
	case 4:
		newDirection = NorthWest
	case 5:
		newDirection = NorthEast
	case 6:
		newDirection = SouthWest
	case 7:
		newDirection = SouthEast
	}

	if !allowOpposite && newDirection.opposite() == currentDirection {
		return currentDirection
	}
	return newDirection
}

func (d Direction) opposite() Direction {
	return Direction{varX: -d.varX, varY: -d.varY}
}
