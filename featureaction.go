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
