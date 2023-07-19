package main

import (
	"bufio"
	"encoding/gob"
	"os"
)



func saveState(filename string) {

	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	writer := bufio.NewWriter(f)

	enc := gob.NewEncoder(writer)

	for i:=0; i<len(persist); i++ {
		err = enc.Encode(persist[i])
		handle(err)
	}

	writer.Flush()

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

	dec := gob.NewDecoder(bufio.NewReader(f))

	for i:=0; i<len(persist); i++ {
		err = dec.Decode(persist[i])
		handle(err)
	}

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

func handle(e error) {
	if e != nil {
		panic(e)
	}
}
