package main

var (
	itemActions map[string]func(i *item)
)

func init() {
	itemActions = make(map[string]func(i *item))
	itemActions["eat"] = eatItem
	itemActions["drop"] = dropItem
}

func eatItem(i *item) {
	p.items.remove(i)
	messages.push("You ate " + i.Prefix + " " + i.Name)
}

func dropItem(i *item) {

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
	messages.push("You dropped " + i.Prefix + " " + i.Name)

}
