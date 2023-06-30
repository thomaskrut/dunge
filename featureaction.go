package main

func open() {

	action := func(dir direction) bool {
		newPosition := p.getPosition()
		newPosition.move(dir)
		if f, ok := featuresOnMap[newPosition]; ok {
			if f.name == "door" && f.obstacle {
				f.obstacle = false
				d.grid[newPosition.x][newPosition.y] = empty
				alterAreaVisibility(&d, p.position, lit, p.lightsource)
				f.char = "-"
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
		newPosition := p.getPosition()
		newPosition.move(dir)
		if f, ok := featuresOnMap[newPosition]; ok {
			if f.name == "door" && !f.obstacle {
				f.obstacle = true
				alterAreaVisibility(&d, p.position, visited, p.lightsource)
				d.grid[newPosition.x][newPosition.y] = wall
				alterAreaVisibility(&d, p.position, lit, p.lightsource)
				f.char = "+"
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

			if d.grid[currentPosition.x][currentPosition.y]&lit != lit {
				break
			}

			if f, ok := featuresOnMap[currentPosition]; ok {
				messages.push("You see a " + f.name, gameplay)
				arrows.push(point{currentPosition.x, currentPosition.y+1})
			}

			if m, ok := monstersOnMap[currentPosition]; ok {
				messages.push("You see a " + m.Name, gameplay)
				arrows.push(point{currentPosition.x, currentPosition.y+1})
			}

			if i, ok := itemsOnMap[currentPosition]; ok {
				messages.push("You see " + i.Prefix + " " + i.Name, gameplay)
				arrows.push(point{currentPosition.x, currentPosition.y+1})
			}

		}

		if len(messages.messageQueue) == 0 {
			messages.push("You don't see anything of interest in that direction", gameplay)
		}
		return false
	}
	messages.push("Which direction?", newDirSelect(action))

}