package main

func open() {

	action := func(dir direction) bool {
		newPosition := p.getPosition()
		newPosition.move(dir)
		if f, ok := featuresOnMap[newPosition]; ok {
			if f.name == "door" && f.obstacle {
				f.obstacle = false
				d.grid[newPosition.x][newPosition.y] = empty
				p.attemptMove(None)
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
			}

			if m, ok := monstersOnMap[currentPosition]; ok {
				messages.push("You see a " + m.Name, gameplay)
			}

			if i, ok := itemsOnMap[currentPosition]; ok {
				messages.push("You see " + i.Prefix + " " + i.Name, gameplay)
			}

		}

		if len(messages.messageQueue) == 0 {
			messages.push("You don't see anything of interest in that direction", gameplay)
		}
		return false
	}
	messages.push("Which direction?", newDirSelect(action))

}