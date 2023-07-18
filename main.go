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
	charmap characterMapper
	dungeon dungeonMap
	p       player

	currentState keyProcessor
	messages     messagePrompt
	gameplay     gamePlay

	monstersOnMap map[point]*monster
	itemsOnMap    map[point][]*item
	featuresOnMap map[point]*feature

	arrows arrowQueue

	validKeyPressed bool

	gridOverlay  []string
	menuItems    []*item
	selectedItem int
	turn         int
)

const (
	obstacle byte = 0
	empty    byte = 1 << iota
	visited
	lit
	room
)

func init() {

	monstersOnMap = make(map[point]*monster)
	itemsOnMap = make(map[point][]*item)
	featuresOnMap = make(map[point]*feature)
	charmap = initCharMap()

	savedStateLoaded := loadState()

	if !savedStateLoaded {
		dungeon = newDungeon(width, height)
		generateDungeon()
		p = newPlayer('@')
		p.setPosition(dungeon.getEmptyPoint())
		turn = 0
		generateItems(readItemsTemplate(), 50)
	}

	generateMonsters(readMonsterTemplate(), 50)

}

func moveMonsters() {

	for i := 0; i < p.Speed; i++ {

		for i, m := range monstersOnMap {

			if m.readyToMove() {

				if items, ok := itemsOnMap[m.position]; ok && m.CarriesItems && randomNumber(20) > m.Speed {
					if dungeon.read(m.position)&lit == lit {
						messages.push("The "+m.Name+" picked up "+items[len(items)-1].Prefix+" "+items[len(items)-1].Name, gameplay)
					}
					m.items.add(items[len(items)-1])
					itemsOnMap[m.position] = itemsOnMap[m.position][:len(itemsOnMap[m.position])-1]
					if len(itemsOnMap[m.position]) == 0 {
						delete(itemsOnMap, m.position)
					}
					continue
				}

				var newDirection direction
				newDirection.connect(m.getPosition(), p.getPosition())
				if !m.MovesDiagonally {
					newDirection.toNonDiagonal()
				}
				for i := 0; !m.attemptMove(newDirection) && i < 10; i++ {
					newDirection = randomDirection(newDirection, false, m.MovesDiagonally)
				}
				delete(monstersOnMap, i)
				monstersOnMap[m.position] = m
			}

		}

	}
}

func checkForItems() {

	if i, ok := itemsOnMap[p.Position]; ok {
		if len(i) == 1 {
			messages.push("There is "+i[0].Prefix+" "+i[0].Name+" here, press 5 to pick up", gameplay)
		} else if len(i) > 1 {
			messages.push("There are some things there, press 5 to examine", gameplay)
		}

	}

}

func generateMonsters(list monsterList, numberOfIterations int) {

	for i := 0; i < numberOfIterations; i++ {

		rand := randomNumber(1000)

		for _, m := range list.Monsters {
			if rand < m.Prob {
				newMonster := m
				newMonster.setPosition(dungeon.getEmptyPoint())
				newMonster.items = newInventory()
				newMonster.speedCounter = newMonster.Speed
				monstersOnMap[newMonster.position] = &newMonster
			}
		}
	}
}

func generateOverlay(menu bool, verb string) {

	gridOverlay = nil
	menuItems = nil
	cursor := "| "

	if verb == "pick up" {

		for _, item := range itemsOnMap[p.Position] {
			itemToAdd := item
			menuItems = append(menuItems, itemToAdd)
		}

	} else {

		if p.Items.count() == 0 {
			messages.push("Inventory empty", gameplay)
			currentState = gameplay
			return
		}

		for item := range p.Items.all() {
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
				newItem.setPosition(dungeon.getEmptyPoint())
				itemsOnMap[newItem.position] = append(itemsOnMap[newItem.position], &newItem)
			}
		}
	}

}

func generateDungeon() {

	var previousRoom point
	var nextRoom point
	var err error

	for {
		previousRoom, err = dungeon.createRandomRoom(dungeon.getRandomPoint(), 20, 20)
		if err != nil {
			continue
		}
		break
	}

	for i := 0; i < 10; i++ {

		for {
			nextRoom, err = dungeon.createRandomRoom(dungeon.getRandomPoint(), 20, 20)
			if err != nil {
				continue
			}
			break
		}

		dungeon.connectWithCorridor(previousRoom, nextRoom)
		previousRoom = nextRoom

	}

	dungeon.generateDoors((width + height) / 10)

}

func printDungeon() {
	gridToPrint := render(&dungeon, p, &arrows, gridOverlay, 60, 40, monstersOnMap, itemsOnMap, featuresOnMap)
	//gridToPrint := renderAll(&d, p, &arrows, gridOverlay, monstersOnMap, itemsOnMap, featuresOnMap)
	fmt.Println()
	fmt.Println(string(gridToPrint))
}

func printStats() {
	fmt.Println("HP:", p.Hp, "Turn:", turn)
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
