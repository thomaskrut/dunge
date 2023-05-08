package main

const (
	northKey = 56
	southKey = 50
	eastKey  = 54
	westKey  = 52
	northWestKey = 55
	northEastKey = 57
	southWestKey = 49
	southEastKey = 51
	restKey = 53
)

type keyProcessor interface {
	keyPressed(char rune)
}

type gameState struct {

}

func (g gameState) keyPressed(char rune) {
	switch char {
	case northKey:
		p.move(North)
	case southKey:
		p.move(South)
	case eastKey:
		p.move(East)
	case westKey:
		p.move(West)
	case northWestKey:
		p.move(NorthWest)
	case northEastKey:
		p.move(NorthEast)
	case southWestKey:
		p.move(SouthWest)
	case southEastKey:
		p.move(SouthEast)
	case restKey:
		p.move(None)
	
}
}