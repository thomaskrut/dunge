package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type itemList struct {
	Items []item `json:"items"`
}

type item struct {
	position Point
	Name     string `json:"name"`
	Prefix   string `json:"prefix"`
	Char     string `json:"char"`
	Prob     int    `json:"prob"`
	Value	 int    `json:"value"`
	Verb	 string `json:"verb"`
}

func (i *item) setPosition(p Point) {
	i.position = p
}

func (i item) getPosition() Point {
	return i.position
}

func (i item) getChar() rune {
	return rune(i.Char[0])
}

func readItemsTemplate() itemList {

	data, err := os.ReadFile("items.json")
	if err != nil {
		panic(err)
	}

	fileAsString := string(data)

	template := itemList{}

	if err := json.Unmarshal([]byte(fileAsString), &template); err != nil {
		panic(err)
	}

	fmt.Println(template)

	return template

}