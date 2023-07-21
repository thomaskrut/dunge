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

	if p.Position.getPossibleDirections(level)[dir] {
		destination := p.Position
		destination.move(dir)

		if m, ok := level.Monsters[destination]; ok {
			p.attack(m)
			return true
		}

		alterAreaVisibility(p.Position, visited, p.Lightsource)

		if !p.InRoom && level.read(destination)&room == room {
			scanRoom(destination, lit)
			p.InRoom = true
		}

		if p.InRoom && level.read(destination)&room == room {
			setRoomState(lit)
		}

		if p.InRoom && level.read(destination)&room != room {
			setRoomState(visited)
			p.CurrentRoom.clear()
			p.InRoom = false
		}

		p.Position.move(dir)

		alterAreaVisibility(p.Position, lit, p.Lightsource)
		fmt.Println(level.read(p.Position))
		return true
	}
	return false
}

func (p *player) pickUpItem() {

	if i, ok := level.Items[p.Position]; ok && len(i) == 1 {
		p.Items.add(i[0])
		delete(level.Items, p.Position)
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

		if level.read(currentPoint)&room == room {
			level.write(currentPoint, empty|room|newState)
			alterAreaVisibility(currentPoint, newState, currentDepth-1)
		} else if level.read(currentPoint)&empty == empty {
			level.write(currentPoint, empty|newState)
			alterAreaVisibility(currentPoint, newState, currentDepth-1)
		} else {
			level.write(currentPoint, newState)
		}
	}
}

func setRoomState(newState byte) {
	for _, p := range p.CurrentRoom.Points {
		if level.read(p)&room == room {
			level.write(p, empty|room|newState)
		} else {
			level.write(p, empty|newState)
		}

	}
}

func scanRoom(pos point, state byte) {

	level.write(pos, empty|room|state)
	for _, dir := range getAllDirections() {
		newPoint := pos
		newPoint.move(dir)
		if level.read(newPoint)&room == room && level.read(newPoint)&state != state {
			p.CurrentRoom.add(newPoint)
			scanRoom(newPoint, state)
		} else if level.read(newPoint)&empty == empty {
			p.CurrentRoom.add(newPoint)
		} else if level.read(newPoint) == obstacle {
			level.write(newPoint, obstacle|visited)
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
