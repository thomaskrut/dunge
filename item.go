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
	Position point
	Name     string   `json:"name"`
	Prefix   string   `json:"prefix"`
	Char     string   `json:"char"`
	Prob     int      `json:"prob"`
	Value    int      `json:"value"`
	Weight   int      `json:"weight"`
	Verbs    []string `json:"verbs"`
}

func (i *item) setPosition(p point) {
	i.Position = p
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
