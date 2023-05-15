package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

const (
	width  = 100
	height = 40
)

var (
	charmap          characterMapper
	d                dungeon
	p                player
	currentState     keyProcessor
	messages         messagePrompt
	gameplay         gamePlay
	rooms            []point
	monsterTemplates monsterList
	itemTemplates    itemList
	activeMonsters   map[point]*monster
	activeItems      map[point]*item
	validKeyPressed  bool
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
	currentState = gameplay

	generateMonsters(readMonsterTemplate(), 10)
	generateItems(readItemsTemplate(), 10)

}

func moveMonsters() {

	for i, m := range activeMonsters {

		if m.moveCounter() >= 1 {
			var newDirection direction
			newDirection.connect(m.getPosition(), p.getPosition())
			if !m.Movesdiagonally {
				newDirection.toNonDiagonal()
			}
			for i := 0; !m.move(newDirection) && i < 10; i++ {
				newDirection = randomDirection(newDirection, false, m.Movesdiagonally)
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

func showInventory() {
	fmt.Println(p.inventory)
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
	for i := 0; i < 4; i++ {
		for {
			roomCenterPos, err := d.createRandomRoom(getRandomPoint(&d), 20, 20)

			if err != nil {
				continue
			}
			rooms = append(rooms, roomCenterPos)
			break
		}
	}

	for index, value := range rooms {
		if index < len(rooms)-1 {
			connectWithCorridor(&d, value, rooms[index+1])
		}
	}
	//connectWithCorridor(&d, getEmptyPoint(&d), getEmptyPoint(&d))
}

func main() {

	p.setPosition(getEmptyPoint(&d))
	p.move(None)

	for {

		validKeyPressed = false

		grindToPrint := render(&d, p, activeMonsters, activeItems)

		fmt.Println(string(grindToPrint))
		fmt.Println("HP:", p.hp)

		switch {

		case len(messages.messageQueue) == 1:
			fmt.Print(messages.pop())
			currentState = gameplay

		case len(messages.messageQueue) > 1:
			fmt.Print(messages.pop())
			fmt.Print(" (press space for more...)")
			currentState = messages

		}

		for !validKeyPressed {
			char, _, err := keyboard.GetSingleKey()
			if err != nil {
				panic(err)
			}
			validKeyPressed = currentState.processKey(char)
		}

	}

}
