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
	save("~player.sav", p)
	save("~map.sav", dungeon)
	save("~items.sav", itemsOnMap)
}

func saveMap(f *os.File) {

	f.Write([]byte{byte(dungeon.Width)})
	f.Write([]byte{byte(dungeon.Height)})

	for _, slice := range dungeon.Grid {
		f.Write(slice)
	}
}

func save(path string, source interface{}) {

	if fileExists(path) {
		os.Remove(path)
	}

	f, err := os.Create(path)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	writer := bufio.NewWriter(f)

	enc := gob.NewEncoder(writer)

	err = enc.Encode(source)

	writer.Flush()

	if err != nil {
		panic(err)
	}

}
