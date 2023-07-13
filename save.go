package main

import "os"

func saveMap(filename string) {

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	f.WriteString("map")
	
	for _, slice := range dungeon.grid {
		f.WriteString("row:")
		f.Write(slice)
	}
	//f.Write(dungeon.grid)
}

func savePlayer(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	
	f.WriteString("player")
	slice := []byte{byte(p.position.x), byte(p.position.y)}
	f.Write(slice)

}