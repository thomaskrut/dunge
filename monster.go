package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type monsterList struct {
	Monsters []monster `json:"monsters"`
}

type monster struct {
	moveCounterValue float32
	items            inventory
	position         point

	Char            string   `json:"char"`
	Name            string   `json:"name"`
	Prob            int      `json:"prob"`
	AttackVerbs     []string `json:"attack"`
	Str             int      `json:"str"`
	Hp              int      `json:"hp"`
	Moves           bool     `json:"moves"`
	Aggressive      bool     `json:"aggressive"`
	Speed           float32  `json:"speed"`
	MovesDiagonally bool     `json:"movesdiagonally"`
	CarriesItems    bool     `json:"carriesitems"`
}

func (m *monster) takeDamage(damage int) {
	m.Hp -= damage
	if m.Hp <= 0 {
		messages.push("You killed the " + m.Name, gameplay)
		if m.items.size() > 0 {
			m.dropAllItems()
			messages.push("The " + m.Name + " scattered its belongings on the floor", gameplay)
		}
		delete(monstersOnMap, m.position)
	}
}

func (m *monster) dropAllItems() {

	for item := range m.items.all() {
		currentItem := item
		newPosition := m.getPosition()
		for itemsOnMap[newPosition] != nil {

			dir := randomDirection(None, true, true)
			if newPosition.getPossibleDirections(&d)[dir] {
				newPosition.move(dir)
			}
		}
		currentItem.setPosition(newPosition)
		itemsOnMap[currentItem.position] = currentItem
	}
	m.items.clear()
}

func (m *monster) attack(p *player) {
	if m.Hp > 0 {
		messages.push("The " + m.Name + " " + m.AttackVerbs[randomNumber(len(m.AttackVerbs))] + " you", gameplay)
		p.takeDamage(m.Str)
	}
}

func (m *monster) moveCounter() float32 {
	if m.moveCounterValue >= 1 {
		m.moveCounterValue -= 1
	}

	m.moveCounterValue += m.Speed

	return m.moveCounterValue
}

func (m monster) getPosition() point {
	return m.position
}

func (m *monster) setPosition(p point) {
	m.position = p
}

func (m monster) getChar() rune {
	return rune(m.Char[0])
}

func (m *monster) move(dir direction) bool {

	

	if m.position.getPossibleDirections(&d)[dir] {

		newPoint := m.position
		newPoint.move(dir)
		if newPoint == p.position {
			m.attack(&p)
			return true
		}
		for _, m := range monstersOnMap {
			if m.position == newPoint {
				return false
			}
		}

		m.position.move(dir)

		return true
	}
	return false
}

func readMonsterTemplate() monsterList {

	data, err := os.ReadFile("monsters.json")
	if err != nil {
		panic(err)
	}

	fileAsString := string(data)

	template := monsterList{}

	if err := json.Unmarshal([]byte(fileAsString), &template); err != nil {
		panic(err)
	}

	fmt.Println(template)

	return template

}
