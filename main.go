package main

import "fmt"

const (
	width  = 100
	height = 80
)

var (
	charmap = newCharMap()
	d       = newDungeon(width, height)
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
		if (index < len(rooms) - 1) {
			connectWithCorridor(&d, value, rooms[index + 1])
		}
	}
	connectWithCorridor(&d, getEmptyPoint(&d), getEmptyPoint(&d))

	d.print(charmap)
	fmt.Println(rooms)

}
