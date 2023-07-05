package main

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
	messages.push("You picked up "+i.Prefix+" "+i.Name, gameplay)
	return gameplay
}

func throwItem(i *item) keyProcessor {

	action := func(dir direction) bool {

		maxSteps := 1000 / i.Weight

		newPosition := p.position

		for count := 0; count <= maxSteps; count++ {
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
