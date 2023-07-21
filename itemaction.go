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
	p.Items.add(i)
	level.Items[p.Position] = append(level.Items[p.Position][:selectedItem], level.Items[p.Position][selectedItem+1:]...)
	message := "You picked up " + i.Prefix + " " + i.Name
	if len(level.Items[p.Position]) == 1 {
		message = message + ". There is " + level.Items[p.Position][0].Prefix + " " + level.Items[p.Position][0].Name + " here, press 5 to pick up"
	}
	if len(level.Items[p.Position]) > 1 {
		message = message + ". There are " + strconv.Itoa(len(level.Items[p.Position])) + " more things here, press 5 to examine"
	}
	messages.push(message, gameplay)
	return gameplay
}

func throwItem(i *item) keyProcessor {

	action := func(dir direction) bool {

		distance := 1000 / i.Weight

		newPosition := p.Position

		for count := 0; count <= distance; count++ {
			if newPosition.getPossibleDirections(level)[dir] {
				newPosition.move(dir)

				if level.Monsters[newPosition] != nil {
					level.Monsters[newPosition].takeDamage(i.Weight / 100)
					messages.push("You hit the "+level.Monsters[newPosition].Name+" with "+i.Prefix+" "+i.Name, gameplay)
					break
				}
			} else {
				break
			}
		}

		i.setPosition(newPosition)
		level.Items[newPosition] = append(level.Items[newPosition], i)
		p.Items.remove(i)
		return true
	}

	messages.push("Which direction?", newDirSelect(action))
	gridOverlay = nil
	return messages
}

func eatItem(i *item) keyProcessor {
	p.Items.remove(i)
	messages.push("You ate "+i.Prefix+" "+i.Name, gameplay)
	return gameplay
}

func dropItem(i *item) keyProcessor {

	newPosition := p.getPosition()

	i.setPosition(newPosition)
	level.Items[i.Position] = append(level.Items[i.Position], i)
	p.Items.remove(i)
	messages.push("You dropped "+i.Prefix+" "+i.Name, gameplay)
	return gameplay

}
