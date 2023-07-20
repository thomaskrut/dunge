package main

func useStairs() {

	if f, ok := featuresOnMap[p.Position]; ok {
		switch f.Name {
		case "upstair":
			generateLevel(currentLevel - 1)
			p.attemptMove(None)
			currentState.processTurn()
		case "downstair":
			generateLevel(currentLevel + 1)
			p.attemptMove(None)
			currentState.processTurn()
		}
	} else {
		messages.push("There are no stairs here", gameplay)
	}

}

func open() {

	action := func(dir direction) bool {
		newPosition := p.Position
		newPosition.move(dir)
		if f, ok := featuresOnMap[newPosition]; ok {
			if f.Name == "door" && f.Closed {
				f.Closed = false
				f.Char = "-"
				f.Description = "an open door"
				dungeon.write(newPosition, empty)
				alterAreaVisibility(p.Position, lit, p.Lightsource)
				messages.push("You opened the door", gameplay)
				return true
			} else {
				messages.push("You can't open that", gameplay)
				return false
			}
		} else {
			messages.push("There is nothing to open there", gameplay)
			return false
		}
	}
	messages.push("Which direction?", newDirSelect(action))

}

func close() {

	action := func(dir direction) bool {
		newPosition := p.Position
		newPosition.move(dir)
		if f, ok := featuresOnMap[newPosition]; ok {
			if f.Name == "door" && !f.Closed {
				f.Closed = true
				f.Char = "+"
				f.Description = "a closed door"
				alterAreaVisibility(p.Position, visited, p.Lightsource)
				dungeon.write(newPosition, obstacle)
				alterAreaVisibility(p.Position, lit, p.Lightsource)
				messages.push("You closed the door", gameplay)
				return true
			} else {
				messages.push("You can't close that", gameplay)
				return false
			}
		} else {
			messages.push("There is nothing to close there", gameplay)
			return false
		}
	}
	messages.push("Which direction?", newDirSelect(action))

}

func look() {

	action := func(dir direction) bool {

		currentPosition := p.getPosition()

		for {

			currentPosition.move(dir)

			if dungeon.read(currentPosition)&lit != lit {
				break
			}

			if f, ok := featuresOnMap[currentPosition]; ok {
				messages.push("You see "+f.Description, gameplay)
				arrows.push(point{currentPosition.X, currentPosition.Y + 1})
			}

			if m, ok := monstersOnMap[currentPosition]; ok {
				messages.push("You see a "+m.Name, gameplay)
				arrows.push(point{currentPosition.X, currentPosition.Y + 1})
			}

			if i, ok := itemsOnMap[currentPosition]; ok {
				if len(i) == 1 {
					messages.push("You see "+i[0].Prefix+" "+i[0].Name, gameplay)
				} else if len(i) > 1 {
					messages.push("You see "+i[len(i)-1].Prefix+" "+i[len(i)-1].Name+" with some other things underneath", gameplay)
				}

				arrows.push(point{currentPosition.X, currentPosition.Y + 1})
			}

		}

		if len(messages.messageQueue) == 0 {
			messages.push("You don't see anything of interest in that direction", gameplay)
		}
		return false
	}
	messages.push("Which direction?", newDirSelect(action))

}
