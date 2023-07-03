package main

func open() {

	action := func(dir direction) bool {
		newPosition := p.position
		newPosition.move(dir)
		if f, ok := featuresOnMap[newPosition]; ok {
			if f.name == "door" && f.obstacle {
				f.obstacle = false
				f.char = "-"
				dungeon.write(newPosition, empty)
				alterAreaVisibility(p.position, lit, p.lightsource)
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
		newPosition := p.position
		newPosition.move(dir)
		if f, ok := featuresOnMap[newPosition]; ok {
			if f.name == "door" && !f.obstacle {
				f.obstacle = true
				f.char = "+"
				alterAreaVisibility(p.position, visited, p.lightsource)
				dungeon.write(newPosition, obstacle)
				alterAreaVisibility(p.position, lit, p.lightsource)
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
				messages.push("You see a "+f.name, gameplay)
				arrows.push(point{currentPosition.x, currentPosition.y + 1})
			}

			if m, ok := monstersOnMap[currentPosition]; ok {
				messages.push("You see a "+m.Name, gameplay)
				arrows.push(point{currentPosition.x, currentPosition.y + 1})
			}

			if i, ok := itemsOnMap[currentPosition]; ok {
				messages.push("You see "+i.Prefix+" "+i.Name, gameplay)
				arrows.push(point{currentPosition.x, currentPosition.y + 1})
			}

		}

		if len(messages.messageQueue) == 0 {
			messages.push("You don't see anything of interest in that direction", gameplay)
		}
		return false
	}
	messages.push("Which direction?", newDirSelect(action))

}
