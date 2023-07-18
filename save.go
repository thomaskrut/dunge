package main

import (
	"bufio"
	"encoding/gob"
	"os"
)

func saveState() {

	save("~player.sav", p)
	save("~map.sav", dungeon)
	save("~items.sav", itemsOnMap)
	save("~monsters.sav", monstersOnMap)
	save("~features.sav", featuresOnMap)
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
