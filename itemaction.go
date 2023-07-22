package main

import "strconv"

var (
	itemActions map[string]func(i *item) keyProcessor
)

func init() {
	itemActions = make(map[string]func(i *item) keyProcessor)
	itemActions["eat"] = eatItem
	itemActions["drop"] = dropItem
	itemActions["throw"] = throwItem
	itemActions["pick up"] = pickUpItem
	itemActions["wield or wear"] = wearItem
}

func wearItem(i *item) keyProcessor {
	return gameplay
}

func pickUpItem(i *item) keyProcessor {
	pl.Items.add(i)
	lvl.Items[pl.Position] = append(lvl.Items[pl.Position][:gridOverlay.selection], lvl.Items[pl.Position][gridOverlay.selection+1:]...)
	message := "You picked up " + i.Prefix + " " + i.Name
	if len(lvl.Items[pl.Position]) == 1 {
		message = message + ". There is " + lvl.Items[pl.Position][0].Prefix + " " + lvl.Items[pl.Position][0].Name + " here, press 5 to pick up"
	}
	if len(lvl.Items[pl.Position]) > 1 {
		message = message + ". There are " + strconv.Itoa(len(lvl.Items[pl.Position])) + " more things here, press 5 to examine"
	}
	messages.push(message, gameplay)
	return gameplay
}

func throwItem(i *item) keyProcessor {

	action := func(dir direction) bool {

		distance := 1000 / i.Weight

		newPosition := pl.Position

		for count := 0; count <= distance; count++ {
			if newPosition.getPossibleDirections(lvl)[dir] {
				newPosition.move(dir)

				if lvl.Monsters[newPosition] != nil {
					lvl.Monsters[newPosition].takeDamage(i.Weight / 100)
					messages.push("You hit the "+lvl.Monsters[newPosition].Name+" with "+i.Prefix+" "+i.Name, gameplay)
					break
				}
			} else {
				break
			}
		}

		i.setPosition(newPosition)
		lvl.Items[newPosition] = append(lvl.Items[newPosition], i)
		pl.Items.remove(i)
		return true
	}

	messages.push("Which direction?", newDirSelect(action))
	gridOverlay.clear()
	return messages
}

func eatItem(i *item) keyProcessor {
	pl.Items.remove(i)
	messages.push("You ate "+i.Prefix+" "+i.Name, gameplay)
	return gameplay
}

func dropItem(i *item) keyProcessor {

	newPosition := pl.getPosition()

	i.setPosition(newPosition)
	lvl.Items[i.Position] = append(lvl.Items[i.Position], i)
	pl.Items.remove(i)
	messages.push("You dropped "+i.Prefix+" "+i.Name, gameplay)
	return gameplay

}
