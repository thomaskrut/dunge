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

	check(err)

	defer f.Close()

	writer := bufio.NewWriter(f)

	enc := gob.NewEncoder(writer)

	for i:=0; i<len(p.entities); i++ {
		err = enc.Encode(p.entities[i])
		check(err)
	}

	writer.Flush()

}

func (p *persist) loadState(filename string) bool {

	if !fileExists(filename) {
		return false
	}

	monstersOnMap = make(map[point]*monster)
	itemsOnMap = make(map[point][]*item)
	featuresOnMap = make(map[point]*feature)

	f, err := os.Open(filename)

	check(err)

	defer f.Close()

	dec := gob.NewDecoder(bufio.NewReader(f))

	for i:=0; i<len(p.entities); i++ {
		err = dec.Decode(p.entities[i])
		check(err)
	}

	check(err)

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

func check(e error) {
	if e != nil {
		panic(e)
	}
}
