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

	for index, _ := range itemsToDisplay {
		if index == selectedItem {
			currentItem := itemsToDisplay[index]
			p.items.remove(currentItem)
			messages.push("You ate " + currentItem.Prefix + " " + currentItem.Name)
			return
		}
	}
}

func dropItem() {

	for i, currentItem := range itemsToDisplay {
		if i == selectedItem {

			newPosition := p.getPosition()

			for itemsOnMap[newPosition] != nil {
				dir := randomDirection(None, true, true)
				if newPosition.getPossibleDirections(&d)[dir] {
					newPosition.move(dir)
				}
			}

			currentItem.setPosition(newPosition)
			itemsOnMap[currentItem.position] = currentItem
			p.items.remove(currentItem)
			messages.push("You dropped " + currentItem.Prefix + " " + currentItem.Name)
			return
		}
	}
}
