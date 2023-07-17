package main

import "os"

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
	savePlayer(f)
}

func saveMap(f *os.File) {

	f.Write([]byte{byte(dungeon.width)})
	f.Write([]byte{byte(dungeon.height)})

	for _, slice := range dungeon.grid {
		f.Write(slice)
	}
}

func savePlayer(f *os.File) {

	slice := []byte{byte(p.position.x), byte(p.position.y)}
	f.Write(slice)

}
