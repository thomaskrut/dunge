package main

import (
	
)

var (
	charmap = NewCharMap()
	dungeon = NewDungeon(80, 30)
)

func main() {

	charmap.Add(0, '#')
	charmap.Add(1, ' ')
	dungeon.SetPoint(20, 10, 1)
	dungeon.Print(charmap)
	
}	