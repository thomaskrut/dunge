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
}

func pickUpItem(i *item) keyProcessor {
	p.items.add(i)
	itemsOnMap[p.position] = append(itemsOnMap[p.position][:selectedItem], itemsOnMap[p.position][selectedItem+1:]...)
	message := "You picked up "+i.Prefix+" "+i.Name
	if len(itemsOnMap[p.position]) == 1 {
		message = message + ". There is " + itemsOnMap[p.position][0].Prefix + " " + itemsOnMap[p.position][0].Name + " here, press 5 to pick up"
	}
	if len(itemsOnMap[p.position]) > 1 {
		message = message + ". There are " + strconv.Itoa(len(itemsOnMap[p.position])) + " more things here, press 5 to examine"
	}
	messages.push(message, gameplay)
	return gameplay
}

func throwItem(i *item) keyProcessor {

	action := func(dir direction) bool {

		distance := 1000 / i.Weight

		newPosition := p.position

		for count := 0; count <= distance; count++ {
			if newPosition.getPossibleDirections(&dungeon)[dir] {
				newPosition.move(dir)

				if monstersOnMap[newPosition] != nil {
					monstersOnMap[newPosition].takeDamage(i.Weight / 100)
					messages.push("You hit the "+monstersOnMap[newPosition].Name+" with "+i.Prefix+" "+i.Name, gameplay)
					break
				}
			} else {
				break
			}
		}

		i.setPosition(newPosition)
		itemsOnMap[newPosition] = append(itemsOnMap[newPosition], i)
		p.items.remove(i)
		return true
	}

	messages.push("Which direction?", newDirSelect(action))
	return messages
}

func eatItem(i *item) keyProcessor {
	p.items.remove(i)
	messages.push("You ate "+i.Prefix+" "+i.Name, gameplay)
	return gameplay
}

func dropItem(i *item) keyProcessor {

	newPosition := p.getPosition()

	i.setPosition(newPosition)
	itemsOnMap[i.position] = append(itemsOnMap[i.position], i)
	p.items.remove(i)
	messages.push("You dropped "+i.Prefix+" "+i.Name, gameplay)
	return gameplay

}
