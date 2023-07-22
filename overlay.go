package main

import (
	"strconv"
	"sort"
	"fmt"
)

func generateOverlay(menu bool, verb string) {

	gridOverlay = nil
	menuItems = nil
	cursor := "| "

	if verb == "pick up" {

		for _, item := range lvl.Items[pl.Position] {
			itemToAdd := item
			menuItems = append(menuItems, itemToAdd)
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
					menuItems = append(menuItems, itemToAdd)
					break
				}
			}
		}
	}

	sort.SliceStable(menuItems, func(i, j int) bool {
		return menuItems[i].Name < menuItems[j].Name
	})

	sort.SliceStable(lvl.Items[pl.Position], func(i, j int) bool {
		return lvl.Items[pl.Position][i].Name < lvl.Items[pl.Position][j].Name
	})

	if len(menuItems) == 0 {
		messages.push("No items to "+verb, gameplay)
		currentState = gameplay
		return
	}

	gridOverlay = append(gridOverlay, "_______________________________")

	if menu {
		gridOverlay = append(gridOverlay, fmt.Sprintf("%-30s%v", "|Select an item to "+verb+":", " |"))
	} else {
		gridOverlay = append(gridOverlay, fmt.Sprintf("%-30s%v", "|Inventory", " |"))
	}

	gridOverlay = append(gridOverlay, fmt.Sprintf("%-30s%v", "|", " |"))

	for index, item := range menuItems {
		if menu {
			if index == selectedItem {
				cursor = "| > "
			} else {
				cursor = "|   "
			}
		}
		gridOverlay = append(gridOverlay, fmt.Sprintf("%-30s%v", cursor+strconv.Itoa(index)+": "+item.Prefix+" "+item.Name, " |"))
	}

	gridOverlay = append(gridOverlay, "|______________________________|")

	return
}