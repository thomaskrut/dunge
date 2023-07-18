package main

import (
	"bufio"
	"encoding/gob"
	"os"
)

func saveState(filename string) {

	if fileExists(filename) {
		os.Remove(filename)
	}

	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	saveMap(f)
	savePlayer()
}

func saveMap(f *os.File) {

	f.Write([]byte{byte(dungeon.width)})
	f.Write([]byte{byte(dungeon.height)})

	for _, slice := range dungeon.grid {
		f.Write(slice)
	}
}

func savePlayer() {

	if fileExists("player.save") {
		os.Remove("player.save")
	}

	f, err := os.Create("player.save")

	if err != nil {
		panic(err)
	}

	defer f.Close()

	writer := bufio.NewWriter(f)

	enc := gob.NewEncoder(writer)

	err = enc.Encode(p)

	writer.Flush()

	if err != nil {
		panic(err)
	}

	//slice := []byte{byte(p.position.x), byte(p.position.y)}
	//f.Write(slice)

}
