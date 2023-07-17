package main

import (
	"os"
)

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func loadState(filename string) bool {

	if !fileExists(filename) {
		return false
	}

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dimensions := make([]byte, 2)
	f.Read(dimensions)
	dungeon.width = int(dimensions[0])
	dungeon.height = int(dimensions[1])

	grid := make([][]byte, dungeon.width)
	for i := range grid {
		row := make([]byte, dungeon.height)
		f.Read(row)
		grid[i] = row
	}

	dungeon.grid = grid

	playerPosition := make([]byte, 2)
	f.Read(playerPosition)
	p.position.x = int(playerPosition[0])
	p.position.y = int(playerPosition[1])

	return true
}