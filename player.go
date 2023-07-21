package main

import "fmt"

type player struct {
	Position    point
	Char        rune
	Lightsource int
	Strength    int
	Hp          int
	Speed       int
	Items       inventory
	InRoom      bool
	CurrentRoom scannedRoom
}

func newPlayer(char rune) player {
	return player{
		Char:        char,
		Lightsource: 4,
		Strength:    8,
		Hp:          16,
		Speed:       10,
		Items:       newInventory(),
	}
}
func (p *player) attemptMove(dir direction) bool {

	if p.Position.getPossibleDirections(lev)[dir] {
		destination := p.Position
		destination.move(dir)

		if m, ok := lev.Monsters[destination]; ok {
			p.attack(m)
			return true
		}

		alterAreaVisibility(p.Position, visited, p.Lightsource)

		if !p.InRoom && lev.read(destination)&room == room {
			scanRoom(destination, lit)
			p.InRoom = true
		}

		if p.InRoom && lev.read(destination)&room == room {
			setRoomState(lit)
		}

		if p.InRoom && lev.read(destination)&room != room {
			setRoomState(visited)
			p.CurrentRoom.clear()
			p.InRoom = false
		}

		p.Position.move(dir)

		alterAreaVisibility(p.Position, lit, p.Lightsource)
		fmt.Println(lev.read(p.Position))
		return true
	}
	return false
}

func (p *player) pickUpItem() {

	if i, ok := lev.Items[p.Position]; ok && len(i) == 1 {
		p.Items.add(i[0])
		delete(lev.Items, p.Position)
		messages.push("You picked up "+i[0].Prefix+" "+i[0].Name, gameplay)
		currentState.processTurn()
	} else if len(i) > 1 {
		currentState = newItemSelect("pick up")
	} else {
		currentState.processTurn()
	}
}

func (p *player) attack(m *monster) {
	messages.push("You hit the "+m.Name, gameplay)
	m.takeDamage(p.Strength)
}

func (p *player) takeDamage(damage int) {
	p.Hp -= damage
}

func alterAreaVisibility(p point, newState byte, currentDepth int) {
	if currentDepth == 0 {
		return
	}
	for _, dir := range getAllDirections() {
		currentPoint := p
		currentPoint.move(dir)

		if lev.read(currentPoint)&room == room {
			lev.write(currentPoint, empty|room|newState)
			alterAreaVisibility(currentPoint, newState, currentDepth-1)
		} else if lev.read(currentPoint)&empty == empty {
			lev.write(currentPoint, empty|newState)
			alterAreaVisibility(currentPoint, newState, currentDepth-1)
		} else {
			lev.write(currentPoint, newState)
		}
	}
}

func setRoomState(newState byte) {
	for _, p := range pl.CurrentRoom.Points {
		if lev.read(p)&room == room {
			lev.write(p, empty|room|newState)
		} else {
			lev.write(p, empty|newState)
		}

	}
}

func scanRoom(pos point, state byte) {

	lev.write(pos, empty|room|state)
	for _, dir := range getAllDirections() {
		newPoint := pos
		newPoint.move(dir)
		if lev.read(newPoint)&room == room && lev.read(newPoint)&state != state {
			pl.CurrentRoom.add(newPoint)
			scanRoom(newPoint, state)
		} else if lev.read(newPoint)&empty == empty {
			pl.CurrentRoom.add(newPoint)
		} else if lev.read(newPoint) == obstacle {
			lev.write(newPoint, obstacle|visited)
		}
	}
}

func (pl player) getPosition() point {
	return pl.Position
}

func (pl *player) setPosition(p point) {
	pl.Position = p
}

func (pl player) getChar() rune {
	return pl.Char
}
