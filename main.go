package main

import (
	"fmt"
	"strconv"

	"github.com/eiannone/keyboard"
)

const (
	width  = 200
	height = 140
)

var (
	charmap         characterMapper
	d               dungeon
	p               player
	currentState    keyProcessor
	messages        messagePrompt
	gameplay        gamePlay
	activeMonsters  map[point]*monster
	activeItems     map[point]*item
	validKeyPressed bool
	gridOverlay     []string
	itemsToDisplay  []*item
	selectedItem    int
	turn            int
)

const (
	wall  = 0
	empty = 1 << iota
	visited
	lit
)

func init() {

	activeMonsters = make(map[point]*monster)
	activeItems = make(map[point]*item)
	charmap = newCharMap()
	charmap.add(wall, ' ')
	charmap.add(empty, ' ')
	charmap.add(empty|visited, ' ')
	charmap.add(empty|lit, '.')
	charmap.add(visited|wall, 9637)
	charmap.add(lit|wall, 9639)

	d = newDungeon(width, height)
	initDungeon()
	p = newPlayer('@')
	turn = 0
	generateMonsters(readMonsterTemplate(), 0)
	generateItems(readItemsTemplate(), 500)

}

func moveMonsters() {

	for i, m := range activeMonsters {

		if m.moveCounter() >= 1 {
			var newDirection direction
			newDirection.connect(m.getPosition(), p.getPosition())
			if !m.MovesDiagonally {
				newDirection.toNonDiagonal()
			}
			for i := 0; !m.move(newDirection) && i < 10; i++ {
				newDirection = randomDirection(newDirection, false, m.MovesDiagonally)
			}
			delete(activeMonsters, i)
			activeMonsters[m.position] = m
		}

	}
}

func checkForItems() {

	if i, ok := activeItems[p.position]; ok {
		messages.push("There is " + i.Prefix + " " + i.Name + " here, press 5 to pick up")
	}

}

func pickUpItem() {

	if i, ok := activeItems[p.position]; ok {
		p.inventory = append(p.inventory, *i)
		delete(activeItems, p.position)
		messages.push("You picked up " + i.Prefix + " " + i.Name)
	}

}

func itemAction(verb string) {
	switch verb {
	case "drop":
		dropItem()
	case "eat":
		eatItem()
	}
}

func eatItem() {

	for i, currentItem := range itemsToDisplay {
		if i == selectedItem {
			p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
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
			p.inventory = append(p.inventory[:i], p.inventory[i+1:]...)
			messages.push("You dropped " + currentItem.Prefix + " " + currentItem.Name)
			return
		}
	}
}

func generateMonsters(list monsterList, numberOfIterations int) {

	for i := 0; i < numberOfIterations; i++ {

		rand := randomNumber(1000)

		for _, m := range list.Monsters {
			if rand < m.Prob {
				newMonster := m
				newMonster.setPosition(getEmptyPoint(&d))
				activeMonsters[newMonster.position] = &newMonster
			}
		}
	}
}

func generateOverlay(menu bool, verb string) {
	if len(p.inventory) == 0 {
		gridOverlay = append(gridOverlay, "Inventory empty")
		currentState = gameplay
		return
	}
	gridOverlay = nil
	itemsToDisplay = nil
	cursor := ""

	for _, item := range p.inventory {
		for _, v := range item.Verbs {
			if v == verb {
				itemToAdd := item
				itemsToDisplay = append(itemsToDisplay, &itemToAdd)
				break
			}
		}
	}

	if len(itemsToDisplay) == 0 {
		gridOverlay = append(gridOverlay, "No items to "+verb)
		currentState = gameplay
		return
	}

	if menu {
		gridOverlay = append(gridOverlay, "Select item to "+verb+":")
	} else {
		gridOverlay = append(gridOverlay, "Inventory:")
	}

	for index, item := range itemsToDisplay {
		if menu {
			if index == selectedItem {
				cursor = "> "
			} else {
				cursor = "  "
			}
		}
		gridOverlay = append(gridOverlay, cursor+strconv.Itoa(index)+": "+item.Prefix+" "+item.Name)
	}
}

func generateItems(list itemList, numberOfIterations int) {

	for i := 1; i < numberOfIterations; i++ {

		rand := randomNumber(1000)

		for _, i := range list.Items {

			if rand < i.Prob {
				newItem := i
				newItem.setPosition(getEmptyPoint(&d))
				activeItems[newItem.position] = &newItem
			}
		}
	}

}

func initDungeon() {

	var previousRoom point
	var nextRoom point
	var err error

	for {
		previousRoom, err = d.createRandomRoom(getRandomPoint(&d), 20, 20)
		if err != nil {
			continue
		}
		break
	}

	for i := 0; i < 10; i++ {

		for {
			nextRoom, err = d.createRandomRoom(getRandomPoint(&d), 20, 20)
			if err != nil {
				continue
			}
			break
		}

		connectWithCorridor(&d, previousRoom, nextRoom)
		previousRoom = nextRoom

	}

}

func printDungeon() {
	grindToPrint := render(&d, p, gridOverlay, 40, 40, activeMonsters, activeItems)
	fmt.Println(string(grindToPrint))
}

func printStats() {
	fmt.Println("HP:", p.hp, "Turn:", turn)
}

func printMessages() {
	switch {
	case len(messages.messageQueue) == 1:
		fmt.Print(messages.pop())
		currentState = gameplay

	case len(messages.messageQueue) > 1:
		fmt.Print(messages.pop())
		fmt.Print(" (press space for more...)")
		currentState = messages

	}
}

func main() {

	currentState = gameplay

	p.setPosition(getEmptyPoint(&d))
	p.attemptMove(None)

	for {

		validKeyPressed = false

		/*switch cs := currentState.(type) {

		case gamePlay, messagePrompt:
			printDungeon()
			printStats()
			printMessages()

		case itemSelect:
			printStats()
		}*/

		printDungeon()
		printStats()
		printMessages()

		for !validKeyPressed {
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			validKeyPressed = currentState.processKey(char)
		}

	}

}
