package main

const (
	northKey     = 56
	southKey     = 50
	eastKey      = 54
	westKey      = 52
	northWestKey = 55
	northEastKey = 57
	southWestKey = 49
	southEastKey = 51
	restKey      = 53
	downStairs   = 62
	upStairs     = 60
	spacebar     = 0
	quitKey      = 113
	inventoryKey = 105
	dropKey      = 100
	eatKey       = 101
	lookKey      = 108
	openKey      = 111
	closeKey     = 99
	throwKey     = 116
	wieldKey     = 119
)

var (
	keyToDirMap = map[rune]direction{
		northKey:     North,
		southKey:     South,
		eastKey:      East,
		westKey:      West,
		northWestKey: NorthWest,
		northEastKey: NorthEast,
		southWestKey: SouthWest,
		southEastKey: SouthEast,
	}
)

type keyProcessor interface {
	processKey(char rune) (validKey bool)
	processTurn()
}
