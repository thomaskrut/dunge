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

	if p.Position.getPossibleDirections(lvl)[dir] {
		destination := p.Position
		destination.move(dir)

		if m, ok := lvl.Monsters[destination]; ok {
			p.attack(m)
			return true
		}

		alterAreaVisibility(p.Position, visited, p.Lightsource)

		if !p.InRoom && lvl.read(destination)&room == room {
			scanRoom(destination, lit)
			p.InRoom = true
		}

		if p.InRoom && lvl.read(destination)&room == room {
			setRoomState(lit)
		}

		if p.InRoom && lvl.read(destination)&room != room {
			setRoomState(visited)
			p.CurrentRoom.clear()
			p.InRoom = false
		}

		p.Position.move(dir)

		alterAreaVisibility(p.Position, lit, p.Lightsource)
		fmt.Println(lvl.read(p.Position))
		return true
	}
	return false
}

func (p *player) pickUpItem() {

	if i, ok := lvl.Items[p.Position]; ok && len(i) == 1 {
		p.Items.add(i[0])
		delete(lvl.Items, p.Position)
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

		if lvl.read(currentPoint)&room == room {
			lvl.write(currentPoint, empty|room|newState)
			alterAreaVisibility(currentPoint, newState, currentDepth-1)
		} else if lvl.read(currentPoint)&empty == empty {
			lvl.write(currentPoint, empty|newState)
			alterAreaVisibility(currentPoint, newState, currentDepth-1)
		} else {
			lvl.write(currentPoint, newState)
		}
	}
}

func setRoomState(newState byte) {
	for _, p := range pl.CurrentRoom.Points {
		if lvl.read(p)&room == room {
			lvl.write(p, empty|room|newState)
		} else {
			lvl.write(p, empty|newState)
		}

	}
}

func scanRoom(pos point, state byte) {

	lvl.write(pos, empty|room|state)
	for _, dir := range getAllDirections() {
		newPoint := pos
		newPoint.move(dir)
		if lvl.read(newPoint)&room == room && lvl.read(newPoint)&state != state {
			pl.CurrentRoom.add(newPoint)
			scanRoom(newPoint, state)
		} else if lvl.read(newPoint)&empty == empty {
			pl.CurrentRoom.add(newPoint)
		} else if lvl.read(newPoint) == obstacle {
			lvl.write(newPoint, obstacle|visited)
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
