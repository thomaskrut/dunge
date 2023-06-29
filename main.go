package main

import (
	"fmt"
	"sort"
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
	monstersOnMap   map[point]*monster
	itemsOnMap      map[point]*item
	featuresOnMap   map[point]*feature
	validKeyPressed bool
	gridOverlay     []string
	menuItems       []*item
	selectedItem    int
	turn            int
)

const (
	wall  = 0
	empty = 1 << iota
	visited
	lit
	arrow
)

func init() {

	monstersOnMap = make(map[point]*monster)
	itemsOnMap = make(map[point]*item)
	featuresOnMap = make(map[point]*feature)
	charmap = initChapMap()

	d = newDungeon(width, height)
	initDungeon()
	p = newPlayer('@')
	turn = 0
	generateMonsters(readMonsterTemplate(), 30)
	generateItems(readItemsTemplate(), 50)

}

func moveMonsters() {

	for i, m := range monstersOnMap {

		if m.moveCounter() >= 1 {

			if item, ok := itemsOnMap[m.position]; ok && m.CarriesItems {
				if d.grid[m.position.x][m.position.y]&lit == lit {
					messages.push("The "+m.Name+" picked up "+item.Prefix+" "+item.Name, gameplay)
				}
				m.items.add(item)
				delete(itemsOnMap, m.position)
				continue
			}

			var newDirection direction
			newDirection.connect(m.getPosition(), p.getPosition())
			if !m.MovesDiagonally {
				newDirection.toNonDiagonal()
			}
			for i := 0; !m.move(newDirection) && i < 10; i++ {
				newDirection = randomDirection(newDirection, false, m.MovesDiagonally)
			}
			delete(monstersOnMap, i)
			monstersOnMap[m.position] = m
		}

	}
}

func checkForItems() {

	if i, ok := itemsOnMap[p.position]; ok {
		messages.push("There is "+i.Prefix+" "+i.Name+" here, press 5 to pick up", gameplay)
	}

}

func pickUpItem() {

	if i, ok := itemsOnMap[p.position]; ok {
		p.items.add(i)
		delete(itemsOnMap, p.position)
		messages.push("You picked up "+i.Prefix+" "+i.Name, gameplay)
	}

}

func generateMonsters(list monsterList, numberOfIterations int) {

	for i := 0; i < numberOfIterations; i++ {

		rand := randomNumber(1000)

		for _, m := range list.Monsters {
			if rand < m.Prob {
				newMonster := m
				newMonster.setPosition(getEmptyPoint(&d))
				newMonster.items = newInventory()
				monstersOnMap[newMonster.position] = &newMonster
			}
		}
	}
}

func generateOverlay(menu bool, verb string) {
	if p.items.size() == 0 {
		messages.push("Inventory empty", gameplay)
		currentState = gameplay
		return
	}
	gridOverlay = nil
	menuItems = nil
	cursor := "| "

	for item := range p.items.all() {
		for _, v := range item.Verbs {
			if v == verb {
				itemToAdd := item
				menuItems = append(menuItems, itemToAdd)
				break
			}
		}
	}

	sort.SliceStable(menuItems, func(i, j int) bool {
		return menuItems[i].Name < menuItems[j].Name
	})

	if len(menuItems) == 0 {
		messages.push("No items to "+verb, gameplay)
		currentState = gameplay
		return
	}

	longestItemName := 0
	for _, item := range menuItems {
		if len(item.Name) > longestItemName {
			longestItemName = len(item.Name)
		}
	}

	frameTop := ""
	for i := 0; i < longestItemName+8; i++ {
		frameTop += "-"
	}

	gridOverlay = append(gridOverlay, frameTop)

	if menu {
		gridOverlay = append(gridOverlay, "| Select item to "+verb+":")
	} else {
		gridOverlay = append(gridOverlay, "| Inventory:")
	}

	for index, item := range menuItems {
		if menu {
			if index == selectedItem {
				cursor = "| > "
			} else {
				cursor = "|   "
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
				itemsOnMap[newItem.position] = &newItem
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

	d.generateDoors(100)

}

func printDungeon() {
	grindToPrint := render(&d, p, gridOverlay, 60, 40, monstersOnMap, itemsOnMap, featuresOnMap)
	fmt.Println()
	fmt.Println(string(grindToPrint))
}

func printStats() {
	fmt.Println("HP:", p.hp, "Turn:", turn)
}

func printMessages() {
	switch {
	case len(messages.messageQueue) == 1:
		fmt.Print(messages.pop())
		currentState = messages.revertToState

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
