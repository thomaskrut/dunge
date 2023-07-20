package main

import (
	"bufio"
	"encoding/gob"
	"os"
)


type persist struct {
	entities map[int]interface{}
}

func (p *persist) register (entities ...interface{}) {
	p.entities = make(map[int]interface{})
	
	for i, e := range entities {
		p.entities[i] = e
	}

}


func (p *persist) saveState(filename string) {

	f, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	writer := bufio.NewWriter(f)

	enc := gob.NewEncoder(writer)

	for i:=0; i<len(p.entities); i++ {
		err = enc.Encode(p.entities[i])
		handle(err)
	}

	writer.Flush()

}

func (p *persist) loadState(filename string) bool {

	if !fileExists(filename) {
		return false
	}

	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	dec := gob.NewDecoder(bufio.NewReader(f))

	for i:=0; i<len(p.entities); i++ {
		err = dec.Decode(p.entities[i])
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
