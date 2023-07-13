package main

type player struct {
	position    point
	char        rune
	lightsource int
	strength    int
	hp          int
	speed       int
	items       inventory
	inRoom      bool
	currentRoom scannedRoom
}

func newPlayer(char rune) player {
	return player{
		char:        char,
		lightsource: 4,
		strength:    8,
		hp:          16,
		speed:       10,
		items:       newInventory(),
	}
}
func (p *player) attemptMove(dir direction) bool {
	
	if p.position.getPossibleDirections(&dungeon)[dir] {
		destination := p.position
		destination.move(dir)

		if m, ok := monstersOnMap[destination]; ok {
			p.attack(m)
			return true
		}

		alterAreaVisibility(p.position, visited, p.lightsource)

		if !p.inRoom && dungeon.read(destination)&room == room {
			scanRoom(destination, lit)
			p.inRoom = true
		}

		if p.inRoom && dungeon.read(destination)&room == room {
			setRoomState(lit)
		}

		if p.inRoom && dungeon.read(destination)&room != room {
			setRoomState(visited)
			p.currentRoom.clear()
			p.inRoom = false
		}

		p.position.move(dir)

		alterAreaVisibility(p.position, lit, p.lightsource)
		return true
	}
	return false
}

func (p *player) pickUpItem() {

	if i, ok := itemsOnMap[p.position]; ok && len(i) == 1 {
		p.items.add(i[0])
		delete(itemsOnMap, p.position)
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
	m.takeDamage(p.strength)
}

func (p *player) takeDamage(damage int) {
	p.hp -= damage
}

func alterAreaVisibility(p point, newState int, currentDepth int) {
	if currentDepth == 0 {
		return
	}
	for _, dir := range getAllDirections() {
		currentPoint := p
		currentPoint.move(dir)

		if dungeon.read(currentPoint)&room == room {
			dungeon.write(currentPoint, empty|room|newState)
			alterAreaVisibility(currentPoint, newState, currentDepth-1)
		} else if dungeon.read(currentPoint)&empty == empty {
			dungeon.write(currentPoint, empty|newState)
			alterAreaVisibility(currentPoint, newState, currentDepth-1)
		} else {
			dungeon.write(currentPoint, newState)
		}
	}
}

func setRoomState(newState int) {
	for _, p := range p.currentRoom.points {
		if dungeon.grid[p.x][p.y]&room == room {
			dungeon.grid[p.x][p.y] = empty | room | newState
		} else {
			dungeon.grid[p.x][p.y] = empty | newState
		}

	}
}

func scanRoom(pos point, state int) {

	dungeon.grid[pos.x][pos.y] = empty | room | state
	for _, dir := range getAllDirections() {
		newPoint := pos
		newPoint.move(dir)
		if dungeon.grid[newPoint.x][newPoint.y]&room == room && dungeon.grid[newPoint.x][newPoint.y]&state != state {
			p.currentRoom.add(newPoint)
			scanRoom(newPoint, state)
		} else if dungeon.grid[newPoint.x][newPoint.y]&empty == empty {
			p.currentRoom.add(newPoint)
		} else if dungeon.grid[newPoint.x][newPoint.y] == obstacle {
			dungeon.grid[newPoint.x][newPoint.y] = obstacle | visited
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

