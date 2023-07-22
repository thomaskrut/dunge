package main

import (
	"fmt"
	"sort"
	"strconv"
)

type overlay struct {
	menu      []string
	menuItems []*item
	selection int
}

func (o *overlay) selectedItem() *item {
	return o.menuItems[o.selection]
}

func (o *overlay) clear() {
	o.menu = nil
}

func (o *overlay) cursorUp() {
	o.selection--
	if o.selection < 0 {
		o.selection = len(o.menuItems) - 1
	}
}

func (o *overlay) cursorDown() {
	o.selection++
	if o.selection > len(o.menuItems)-1 {
		o.selection = 0
	}
}

func (o *overlay) generate(menu bool, verb string) {

	o.menu = nil
	o.menuItems = nil
	cursor := "| "

	if verb == "pick up" {

		for _, item := range lvl.Items[pl.Position] {
			itemToAdd := item
			o.menuItems = append(o.menuItems, itemToAdd)
		}

	} else {

		if pl.Items.count() == 0 {
			messages.push("Inventory empty", gameplay)
			currentState = gameplay
			return
		}

		for item := range pl.Items.all() {
			for _, v := range item.Verbs {
				if v == verb {
					itemToAdd := item
					o.menuItems = append(o.menuItems, itemToAdd)
					break
				}
			}
		}
	}

	sort.SliceStable(o.menuItems, func(i, j int) bool {
		return o.menuItems[i].Name < o.menuItems[j].Name
	})

	sort.SliceStable(lvl.Items[pl.Position], func(i, j int) bool {
		return lvl.Items[pl.Position][i].Name < lvl.Items[pl.Position][j].Name
	})

	if len(o.menuItems) == 0 {
		messages.push("No items to "+verb, gameplay)
		currentState = gameplay
		return
	}

	o.menu = append(o.menu, "_______________________________")

	if menu {
		o.menu = append(o.menu, fmt.Sprintf("%-30s%v", "|Select an item to "+verb+":", " |"))
	} else {
		o.menu = append(o.menu, fmt.Sprintf("%-30s%v", "|Inventory", " |"))
	}

	o.menu = append(o.menu, fmt.Sprintf("%-30s%v", "|", " |"))

	for index, item := range o.menuItems {
		if menu {
			if index == o.selection {
				cursor = "| > "
			} else {
				cursor = "|   "
			}
		}
		o.menu = append(o.menu, fmt.Sprintf("%-30s%v", cursor+strconv.Itoa(index)+": "+item.Prefix+" "+item.Name, " |"))
	}

	o.menu = append(o.menu, "|______________________________|")

	return
}
