package main

import (
	"fmt"
	"strconv"

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
	rooms            []Point
	monsterTemplates monsterList
	itemTemplates	 itemList
	activeMonsters   []monster
	activeItems	  	 []item
	validKeyPressed  bool
	numberOfItemsFound int
)

const (
	wall  = 0
	empty = 1 << iota
	visited
	lit
)

func init() {

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
	currentState = gamePlay{}

	monsterTemplates = readMonsterTemplate()
	itemTemplates = readItemsTemplate()
	activeMonsters = generateMonsters(10)
	activeItems = generateItems(10)

}

func moveMonsters() {

	for i := range activeMonsters {
		m := &activeMonsters[i]
		if m.moveCounter() >= 1 {
			var newDirection direction
			newDirection.connect(m.getPosition(), p.getPosition())
			if !m.Movesdiagonally {
				newDirection.toNonDiagonal()
			}
			for i:=0; !m.move(newDirection) && i < 10; i++ {
				newDirection = randomDirection(newDirection, false, m.Movesdiagonally)
			}
		}

	}
}

func checkForItems() int {
	count := 0
	var itemsMessage = "You have stumbled upon"
	for i := range activeItems {
		item := activeItems[i]
		if item.getPosition().overlaps(p.getPosition()) {
			count++
			itemsMessage = itemsMessage + ", " + strconv.Itoa(count) + ": " + item.Prefix + " " + item.Name
		}
	}
	if count > 0 {
		itemsMessage = itemsMessage + ". Pick up which? (or press space to continue)"
		messages.addMessage(itemsMessage)
		return count
	}
	return 0
}

func checkMonsterHealth() {
	for i, m := range activeMonsters {
		if m.Hp <= 0 {
			activeMonsters = append(activeMonsters[:i], activeMonsters[i+1:]...)
		}
	}
}

func generateMonsters(numberOfIterations int) []monster {

	var monsterSlice []monster

	for i := 0; i < numberOfIterations; i++ {

		rand := randomNumber(1000)

		for _, m := range monsterTemplates.Monsters {
			if rand < m.Prob {
				newMonster := m
				newMonster.setPosition(getEmptyPoint(&d))
				monsterSlice = append(monsterSlice, newMonster)
			}
		}

	}
	return monsterSlice
}

func generateItems(numberOfIterations int) []item {
	var itemSlice []item

	for i:=1; i < numberOfIterations; i++ {
		rand := randomNumber(1000)
		for _, i := range itemTemplates.Items {

		if rand < i.Prob {
				newItem := i
				newItem2 := i
				newItem.setPosition(getEmptyPoint(&d))
				newItem2.setPosition(newItem.position)
				itemSlice = append(itemSlice, newItem)
				itemSlice = append(itemSlice, newItem2)
			}
		}

	}
	return itemSlice
}

func initDungeon() {
	for i := 0; i < 4; i++ {
		for {
			position, err := d.createRandomRoom(getRandomPoint(&d), 20, 20)

			if err != nil {
				continue
			}
			rooms = append(rooms, position)
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

		if len(messages.messageQueue) == 1 {
			fmt.Print(messages.getOldestMessage())
			messages.deleteOldestMessage()
			currentState = gamePlay{}
		} else if len(messages.messageQueue) > 1 {
			fmt.Print(messages.getOldestMessage())
			messages.deleteOldestMessage()
			fmt.Print(" (press space for more...)")
			currentState = messages
		}

		if numberOfItemsFound > 0 {
			currentState = newPickupState(numberOfItemsFound)
			numberOfItemsFound = 0
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
