package main

import "os"

func saveMap(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("map")

	//f.Write(dungeon.grid)



}