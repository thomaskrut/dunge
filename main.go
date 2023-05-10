package main

import (
	"fmt"
	"math/rand"

	"github.com/eiannone/keyboard"
)

const (
	width  = 100
	height = 20
)

var (
	charmap          characterMapper
	d                dungeon
	p                player
	currentState     keyProcessor
	rooms            []Point
	monsterTemplates monsterList
	activeMonsters   []entity
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
	activeMonsters = generateMonsters(10)

}

func moveMonsters() {
	for _, m := range activeMonsters {
		if m.moveCounter() >= 1 {
			var newDirection direction
			newDirection.directionTowards(m.getPosition(), p.getPosition())
			for !m.move(newDirection) {
				newDirection = randomDirection(newDirection, true, true)
			}
		}

	}
}

func generateMonsters(numberOfIterations int) []entity {

	var monsterSlice []entity

	for i := 0; i < numberOfIterations; i++ {

		rand := rand.Intn(1000)

		for _, m := range monsterTemplates.Monsters {
			if rand < m.Prob {

				newMonster := m
				newMonster.setPosition(getEmptyPoint(&d))
				monsterSlice = append(monsterSlice, &newMonster)
			}
		}

	}
	return monsterSlice
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

		grindToPrint := render(&d, p, activeMonsters...)
		fmt.Println(string(grindToPrint))
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		currentState.processKey(char)

	}

}
