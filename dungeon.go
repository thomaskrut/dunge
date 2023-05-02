package main

import (
	"fmt"
	"math/rand"
)

type Dungeon struct {
	grid [][]int
}

func NewDungeon(width, height int) Dungeon {
	zeroedGrid := make([][]int, height)
	for i := range zeroedGrid {
		zeroedGrid[i] = make([]int, width)
	}
	return Dungeon{grid: zeroedGrid}
}

func (d Dungeon) Print(charmap characterMapper) {
	for _, row := range d.grid {
		for _, cell := range row {
			fmt.Printf("%c", charmap.chars[cell])
		}
		println()
	}
}

func (d *Dungeon) SetPoint(x, y int, value int) {
	d.grid[y][x] = value
}

func (d *Dungeon) GetPoint(x, y int) int {
	return d.grid[y][x]
}

func getEmptyPoint(d Dungeon) (int, int) {
	for {
		x := rand.Intn(len(d.grid))
		y := rand.Intn(len(d.grid[0]))
		if d.grid[x][y] == 1 {
			return x, y
		}
	}
	}
