package main

import (
	"bufio"
	"encoding/gob"
	"os"
)

func loadState() bool {

	if !fileExists("~player.sav") {
		return false
	}

	load("~player.sav", &p)
	load("~map.sav", &dungeon)
	load("~items.sav", &itemsOnMap)
	load("~monsters.sav", &monstersOnMap)
	load("~features.sav", &featuresOnMap)
	return true

}

func load(filename string, target interface{}) bool {
	if !fileExists(filename) {
		return false
	}
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)

	dec := gob.NewDecoder(reader)

	err = dec.Decode(target)

	if err != nil {
		panic(err)
	}

	return true
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
