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
	empty = 1 << iota
	visited
	lit
)

func main() {

	charmap.add(0, 9639) //wall
	charmap.add(1, ' ')
	charmap.add(2, '.')
	charmap.add(4, ':')
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
	connectWithCorridor(&d, getEmptyPoint(&d), getEmptyPoint(&d))

	p = newPlayer()
	p.position = getEmptyPoint(&d)

	currentState = gameState{}

	for {

		grindToPrint := render(&d, &p)

		fmt.Println(string(grindToPrint))
		fmt.Println(p)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		currentState.keyPressed(char)

	}

}
