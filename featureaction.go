package main

func useStairs(direction string) {

	if f, ok := lev.Features[pl.Position]; ok {
		if f.State != direction {
			messages.push("There are no stairs to walk "+direction+" here", gameplay)
			return
		}
		switch direction {
		case "up":
			world.CurrentDepth--
			if l, ok := world.Levels[world.CurrentDepth]; ok {
				setRoomState(visited)
				lev = l
				pl.Position = lev.Downstair
				pl.CurrentRoom.clear()
				pl.InRoom = false
			} else {
				generateLevel(world.CurrentDepth)
			}
			pl.attemptMove(None)
			currentState.processTurn()
		case "down":
			world.CurrentDepth++
			if l, ok := world.Levels[world.CurrentDepth]; ok {
				setRoomState(visited)
				lev = l
				pl.Position = lev.Upstair
				pl.CurrentRoom.clear()
				pl.InRoom = false
			} else {
				generateLevel(world.CurrentDepth)
			}
			pl.attemptMove(None)
			currentState.processTurn()
		}
	} else {
		messages.push("There are no stairs here", gameplay)
	}

}

func open() {

	action := func(dir direction) bool {
		newPosition := pl.Position
		newPosition.move(dir)
		if f, ok := lev.Features[newPosition]; ok {
			if f.Name == "door" && f.State == "closed" {
				f.State = "open"
				f.Char = "-"
				f.Description = "an open door"
				lev.write(newPosition, empty)
				alterAreaVisibility(pl.Position, lit, pl.Lightsource)
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
		newPosition := pl.Position
		newPosition.move(dir)
		if f, ok := lev.Features[newPosition]; ok {
			if f.Name == "door" && f.State == "open" {
				f.State = "closed"
				f.Char = "+"
				f.Description = "a closed door"
				alterAreaVisibility(pl.Position, visited, pl.Lightsource)
				lev.write(newPosition, obstacle)
				alterAreaVisibility(pl.Position, lit, pl.Lightsource)
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

		currentPosition := pl.getPosition()

		for {

			currentPosition.move(dir)

			if lev.read(currentPosition)&lit != lit {
				break
			}

			if f, ok := lev.Features[currentPosition]; ok {
				messages.push("You see "+f.Description, gameplay)
				arrows.push(point{currentPosition.X, currentPosition.Y + 1})
			}

			if m, ok := lev.Monsters[currentPosition]; ok {
				messages.push("You see a "+m.Name, gameplay)
				arrows.push(point{currentPosition.X, currentPosition.Y + 1})
			}

			if i, ok := lev.Items[currentPosition]; ok {
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
