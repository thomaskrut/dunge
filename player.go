package main

import "fmt"

type player struct {
	position    point
	char        rune
	lightsource int
	strength    int
	hp          int
	items       inventory
	inRoom      bool
	currentRoom scannedRoom
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

		if !p.inRoom && d.grid[newPoint.x][newPoint.y]&room == room {
			scanRoom(newPoint, lit)
			p.inRoom = true
		}

		if p.inRoom && d.grid[newPoint.x][newPoint.y]&room == room {
			setRoomState(lit)
		}

		if p.inRoom && d.grid[newPoint.x][newPoint.y]&room != room {
			setRoomState(visited)
			p.currentRoom.clear()
			p.inRoom = false
		}

		p.position.move(dir)

		alterAreaVisibility(&d, p.position, lit, p.lightsource)
		return true
	}
	return false
}

func (p *player) attack(m *monster) {
	messages.push("You hit the "+m.Name, gameplay)
	m.takeDamage(p.strength)
}

func (p *player) takeDamage(damage int) {
	p.hp -= damage
}

func alterAreaVisibility(d *dungeon, p point, state int, currentDepth int) {
	if currentDepth == 0 {
		return
	}
	for _, dir := range getAllDirections() {
		newPoint := p
		newPoint.move(dir)
		if d.grid[newPoint.x][newPoint.y]&room == room {
			d.grid[newPoint.x][newPoint.y] = empty | room | state
			alterAreaVisibility(d, newPoint, state, currentDepth-1)
		} else if d.grid[newPoint.x][newPoint.y]&empty == empty {
			d.grid[newPoint.x][newPoint.y] = empty | state
			alterAreaVisibility(d, newPoint, state, currentDepth-1)
		} else {
			d.grid[newPoint.x][newPoint.y] = state
		}
	}
}

func setRoomState(state int) {
	for _, p := range p.currentRoom.points {
		if d.grid[p.x][p.y]&room == room {
			d.grid[p.x][p.y] = empty | room | state
		} else {
			d.grid[p.x][p.y] = empty | state
		}

	}
}

func scanRoom(pos point, state int) {

	d.grid[pos.x][pos.y] = empty | room | state
	for _, dir := range getAllDirections() {
		fmt.Println("room")
		newPoint := pos
		newPoint.move(dir)
		if d.grid[newPoint.x][newPoint.y]&room == room && d.grid[newPoint.x][newPoint.y]&state != state {
			p.currentRoom.add(newPoint)
			scanRoom(newPoint, state)
		} else if d.grid[newPoint.x][newPoint.y]&empty == empty {
			p.currentRoom.add(newPoint)
		} else if d.grid[newPoint.x][newPoint.y] == 0 {
			d.grid[newPoint.x][newPoint.y] = wall | visited
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
