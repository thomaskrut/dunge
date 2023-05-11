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
	spacebar     = 0
)

type keyProcessor interface {
	processKey(char rune)
	processTurn()
}
