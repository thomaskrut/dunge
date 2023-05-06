package main

import (
	"fmt"
	_ "fmt"

	"github.com/eiannone/keyboard"
)

const (
	width  = 100
	height = 40
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

	charmap.add(0, '#')
	charmap.add(1, ' ')
	var rooms []Point

	for i := 0; i < 10; i++ {
		for {
			position, err := d.createRandomRoom(getRandomPoint(&d), 15, 15)

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

	p.position = getEmptyPoint(&d)

	currentState = gameState{}

	for {
		d.print(charmap, p)
		fmt.Println(p)
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		currentState.keyPressed(char)

	}

}
