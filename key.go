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
		p.position.move(North)
	case southKey:
		p.position.move(South)
	case eastKey:
		p.position.move(East)
	case westKey:
		p.position.move(West)
	case northWestKey:
		p.position.move(NorthWest)
	case northEastKey:
		p.position.move(NorthEast)
	case southWestKey:
		p.position.move(SouthWest)
	case southEastKey:
		p.position.move(SouthEast)
	case restKey:
		p.position.move(None)
	
}
}