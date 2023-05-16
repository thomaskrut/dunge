package main

import (
	"errors"
)

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

func (im *inventorymenu) getItemByNumber(n int) (*item, error) {
	count := 1
	for itemName := range inventoryMenu.items {
		if count == n {
			item, err := im.getItemByName(itemName)
			if err != nil {
				return nil, err
			} else {
				return item, nil
			}
		}
		count++
	}
	return nil, errors.New("no item found")
}

func (im *inventorymenu) getItemByName(itemName string) (*item, error) {
	for _, i := range p.inventory {
		if i.Name == itemName {
			return &i, nil
		}
	}
	return nil, errors.New("no item found")

}
