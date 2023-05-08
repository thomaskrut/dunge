package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
)

const (
	width  = 100
	height = 20
)

var (
	charmap      = newCharMap()
	d            = newDungeon(width, height)
	p            player
	currentState keyProcessor
)

const (
	wall  = 0
	empty = 1 << iota
	visited
	lit
)

func main() {

	charmap.add(wall, ' ')
	charmap.add(empty, ' ')
	charmap.add(empty|visited, ' ')
	charmap.add(empty|lit, '.')
	charmap.add(visited|wall, 9637)
	charmap.add(lit|wall, 9639)

	var rooms []Point

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

	p = newPlayer()
	p.setPosition(getEmptyPoint(&d))
	p.move(None)
	currentState = gameState{}

	for {

		grindToPrint := render(&d, &p)
		fmt.Println(string(grindToPrint))
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		currentState.processKey(char)

	}

}
