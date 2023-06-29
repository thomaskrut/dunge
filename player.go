package main

type player struct {
	position    point
	char        rune
	lightsource int
	strength    int
	hp          int
	items       inventory
}

func (p *player) attemptMove(dir direction) bool {

	if p.position.getPossibleDirections(&d)[dir] {
		newPoint := p.position
		newPoint.move(dir)

		if m, ok := monstersOnMap[newPoint]; ok {
			p.attack(m)
			return true
		}

		

		alterAreaVisibility(&d, p.position, visited, p.lightsource)
		p.position.move(dir)
		alterAreaVisibility(&d, p.position, lit, p.lightsource)
		return true
	}
	return false
}

func (p *player) attack(m *monster) {
	messages.push("You hit the " + m.Name, gameplay)
	m.takeDamage(p.strength)
}

func (p *player) takeDamage(damage int) {
	p.hp -= damage
}

func alterAreaVisibility(d *dungeon, p point, value int, currentDepth int) {
	if currentDepth == 0 {
		return
	}
	for _, dir := range getAllDirections() {
		newPoint := p
		newPoint.move(dir)
		if d.grid[newPoint.x][newPoint.y]&empty == empty {
			d.grid[newPoint.x][newPoint.y] = empty | value
			alterAreaVisibility(d, newPoint, value, currentDepth-1)
		} else {
			d.grid[newPoint.x][newPoint.y] = value
		}
	}

}

func (pl player) getPosition() point {
	return pl.position
}

func (pl *player) setPosition(p point) {
	pl.position = p
}

func (pl player) getChar() rune {
	return pl.char
}

func newPlayer(char rune) player {
	return player{
		char:        char,
		lightsource: 4,
		strength:    8,
		hp:          16,
		items:       newInventory(),
	}
}
