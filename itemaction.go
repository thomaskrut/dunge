package main

var (
	itemActions map[string]func(i *item) keyProcessor
)

func init() {
	itemActions = make(map[string]func(i *item) keyProcessor)
	itemActions["eat"] = eatItem
	itemActions["drop"] = dropItem
	itemActions["throw"] = throwItem
}

func throwItem(i *item) keyProcessor {

	action := func(dir direction) {

		maxSteps := 1000 / i.Weight

		newPosition := p.position

		for count := 0; count <= maxSteps; count++ {
			if newPosition.getPossibleDirections(&d)[dir] {
				newPosition.move(dir)
				
				if monstersOnMap[newPosition] != nil {
					monstersOnMap[newPosition].takeDamage(i.Weight / 100)
					messages.push("You hit the " + monstersOnMap[newPosition].Name + " with " + i.Prefix + " " + i.Name, gameplay)
					break
				}
			} else {
				break
			}
		}

		i.setPosition(newPosition)
		itemsOnMap[newPosition] = i
		p.items.remove(i)
	}
	
	messages.push("Which direction?", newDirSelect(action))

	return newDirSelect(action)
	
}

func eatItem(i *item) keyProcessor {
	p.items.remove(i)
	messages.push("You ate " + i.Prefix + " " + i.Name, gameplay)
	return gameplay
}

func dropItem(i *item) keyProcessor {

	newPosition := p.getPosition()

	for itemsOnMap[newPosition] != nil {
		dir := randomDirection(None, true, true)
		if newPosition.getPossibleDirections(&d)[dir] {
			newPosition.move(dir)
		}
	}

	i.setPosition(newPosition)
	itemsOnMap[i.position] = i
	p.items.remove(i)
	messages.push("You dropped " + i.Prefix + " " + i.Name, gameplay)
	return gameplay

}
