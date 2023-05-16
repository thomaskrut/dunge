package main

type inventorymenu struct {
	items map[string]int
}

func newInventoryMenu() inventorymenu {

	return inventorymenu{items: make(map[string]int)}

}

func (im *inventorymenu) update() {

	im.items = make(map[string]int)

	for _, i := range p.inventory {

		im.items[i.Name]++

	}

}

func (im *inventorymenu) getItemByNumber(n int) *item {
	count := 1
	for itemName, _ := range inventoryMenu.items {
		if count == n {
			return im.getItemByName(itemName)
		}
		count++
	}
	return nil
}

func (im *inventorymenu) getItemByName(itemName string) *item {

	for _, i := range p.inventory {

		if i.Name == itemName {

			return &i

		}

	}

	return nil

}
