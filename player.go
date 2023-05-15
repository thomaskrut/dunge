package main

type player struct {
	position    Point
	char        rune
	lightsource int
	strength    int
	hp          int
	inventory	[]item
}

func (p *player) move(dir direction) bool {
	
	if p.position.getPossibleDirections(&d)[dir] {
		newPoint := p.position
		newPoint.new(dir)

		if m, ok := activeMonsters[newPoint]; ok {
			p.attack(&m)
			return true
		}

		alterAreaVisibility(&d, p.position, visited, p.lightsource)
		p.position.new(dir)
		alterAreaVisibility(&d, p.position, lit, p.lightsource)
		return true
	}
	return false
}

func (p *player) attack(m *monster) {
	m.takeDamage(p.strength)
	messages.addMessage("You hit the " + m.Name)
	if m.Hp <= 0 {
		messages.addMessage("You killed the " + m.Name)
	}
}

func (p *player) takeDamage(damage int) {
	p.hp -= damage
}

func alterAreaVisibility(d *dungeon, p Point, value int, currentDepth int) {
	if currentDepth == 0 {
		return
	}
	for _, dir := range getAllDirections() {
		newPoint := p
		newPoint.new(dir)
		if d.grid[newPoint.x][newPoint.y]&empty == empty {
			d.grid[newPoint.x][newPoint.y] = empty | value
			alterAreaVisibility(d, newPoint, value, currentDepth-1)
		} else {
			d.grid[newPoint.x][newPoint.y] = value
		}
	}

}

func (p player) getPosition() Point {
	return p.position
}

func (p *player) setPosition(point Point) {
	p.position = point
}

func (p player) getChar() rune {
	return p.char
}

func newPlayer(char rune) player {
	return player{
		char:        char,
		lightsource: 4,
		strength:    8,
		hp:          16,
	}
}
