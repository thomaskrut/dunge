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
	lev.Items[pl.Position] = append(lev.Items[pl.Position][:selectedItem], lev.Items[pl.Position][selectedItem+1:]...)
	message := "You picked up " + i.Prefix + " " + i.Name
	if len(lev.Items[pl.Position]) == 1 {
		message = message + ". There is " + lev.Items[pl.Position][0].Prefix + " " + lev.Items[pl.Position][0].Name + " here, press 5 to pick up"
	}
	if len(lev.Items[pl.Position]) > 1 {
		message = message + ". There are " + strconv.Itoa(len(lev.Items[pl.Position])) + " more things here, press 5 to examine"
	}
	messages.push(message, gameplay)
	return gameplay
}

func throwItem(i *item) keyProcessor {

	action := func(dir direction) bool {

		distance := 1000 / i.Weight

		newPosition := pl.Position

		for count := 0; count <= distance; count++ {
			if newPosition.getPossibleDirections(lev)[dir] {
				newPosition.move(dir)

				if lev.Monsters[newPosition] != nil {
					lev.Monsters[newPosition].takeDamage(i.Weight / 100)
					messages.push("You hit the "+lev.Monsters[newPosition].Name+" with "+i.Prefix+" "+i.Name, gameplay)
					break
				}
			} else {
				break
			}
		}

		i.setPosition(newPosition)
		lev.Items[newPosition] = append(lev.Items[newPosition], i)
		pl.Items.remove(i)
		return true
	}

	messages.push("Which direction?", newDirSelect(action))
	gridOverlay = nil
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
	lev.Items[i.Position] = append(lev.Items[i.Position], i)
	pl.Items.remove(i)
	messages.push("You dropped "+i.Prefix+" "+i.Name, gameplay)
	return gameplay

}
