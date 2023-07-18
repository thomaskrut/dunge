package main

var (
	None      = direction{varX: 0, varY: 0}
	North     = direction{varX: 0, varY: -1}
	South     = direction{varX: 0, varY: 1}
	East      = direction{varX: 1, varY: 0}
	West      = direction{varX: -1, varY: 0}
	NorthWest = direction{varX: -1, varY: -1}
	NorthEast = direction{varX: 1, varY: -1}
	SouthWest = direction{varX: -1, varY: 1}
	SouthEast = direction{varX: 1, varY: 1}
)

type direction struct {
	varX, varY int
}

func getAllDirections() []direction {
	return []direction{North, South, East, West, NorthWest, NorthEast, SouthWest, SouthEast, None}
}

func getNonDiagonalDirections() []direction {
	return []direction{North, South, East, West}
}

func (d *direction) toNonDiagonal() {
	r := randomNumber(2)
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

func (dir *direction) connect(from, to point) *direction {
	switch {
	case to.X == from.X && to.Y < from.Y:
		*dir = North
	case to.X == from.X && to.Y > from.Y:
		*dir = South
	case to.X < from.X && to.Y == from.Y:
		*dir = West
	case to.X > from.X && to.Y == from.Y:
		*dir = East
	case to.X > from.X && to.Y < from.Y:
		*dir = NorthEast
	case to.X > from.X && to.Y > from.Y:
		*dir = SouthEast
	case to.X < from.X && to.Y < from.Y:
		*dir = NorthWest
	case to.X < from.X && to.Y > from.Y:
		*dir = SouthWest
	}
	return dir

}

func randomDirection(currentDirection direction, allowOpposite bool, allowDiagonal bool) direction {

	var r int
	var newDirection direction
	if allowDiagonal {
		r = 8
	} else {
		r = 4
	}
	switch randomNumber(r) {
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

func (d direction) opposite() direction {
	return direction{varX: -d.varX, varY: -d.varY}
}
