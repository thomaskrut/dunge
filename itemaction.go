package main

var (
	itemActions map[string]func()
)

func init() {
	itemActions = make(map[string]func())
	itemActions["eat"] = eatItem
	itemActions["drop"] = dropItem
}

func eatItem() {

	for i, currentItem := range itemsToDisplay {
		if i == selectedItem {
			p.items.remove(i)
			messages.push("You ate " + currentItem.Prefix + " " + currentItem.Name)
			return
		}
	}
}

func dropItem() {

	for i, currentItem := range itemsToDisplay {
		if i == selectedItem {

			newPosition := p.getPosition()

			for activeItems[newPosition] != nil {
				dir := randomDirection(None, true, true)
				if newPosition.getPossibleDirections(&d)[dir] {
					newPosition.move(dir)
				}
			}

			currentItem.setPosition(newPosition)
			activeItems[currentItem.position] = currentItem
			p.items.remove(i)
			messages.push("You dropped " + currentItem.Prefix + " " + currentItem.Name)
			return
		}
	}
}