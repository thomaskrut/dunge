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
	SpeedCounter    int
	Items           inventory
	Position        point
	Char            string   `json:"char"`
	Name            string   `json:"name"`
	Prob            int      `json:"prob"`
	AttackVerbs     []string `json:"attack"`
	Str             int      `json:"str"`
	Hp              int      `json:"hp"`
	Speed           int      `json:"speed"`
	Moves           bool     `json:"moves"`
	Aggressive      bool     `json:"aggressive"`
	MovesDiagonally bool     `json:"movesdiagonally"`
	CarriesItems    bool     `json:"carriesitems"`
}

func (m *monster) takeDamage(damage int) {
	m.Hp -= damage
	m.Aggressive = true
	if m.Hp <= 0 {
		messages.push("You killed the "+m.Name, gameplay)
		if m.Items.count() > 0 {
			m.dropAllItems()
			messages.push("The "+m.Name+" scattered its belongings on the floor", gameplay)
		}
		delete(lev.Monsters, m.Position)
	}
}

func (m *monster) dropAllItems() {

	for item := range m.Items.all() {
		currentItem := item
		newPosition := m.Position
		for lev.Items[newPosition] != nil {

			dir := randomDirection(None, true, true)
			if newPosition.getPossibleDirections(lev)[dir] {
				newPosition.move(dir)
			}
		}
		currentItem.setPosition(newPosition)
		lev.Items[currentItem.Position] = append(lev.Items[currentItem.Position], currentItem)
	}
	m.Items.clear()
}

func (m *monster) attack(p *player) {
	if m.Hp > 0 {
		messages.push("The "+m.Name+" "+m.AttackVerbs[randomNumber(len(m.AttackVerbs))]+" you", gameplay)
		p.takeDamage(m.Str)
	}
}

func (m *monster) readyToMove() bool {
	if m.SpeedCounter == 0 {
		m.SpeedCounter = m.Speed
		return true
	}
	m.SpeedCounter--
	return false
}

func (m monster) getPosition() point {
	return m.Position
}

func (m *monster) setPosition(p point) {
	m.Position = p
}

func (m monster) getChar() rune {
	return rune(m.Char[0])
}

func (m *monster) attemptMove(dir direction) bool {

	if m.Position.getPossibleDirections(lev)[dir] {

		newPoint := m.Position
		newPoint.move(dir)
		if newPoint == pl.Position {
			if m.Aggressive {
				m.attack(&pl)
			}
			return true
		}
		for _, m := range lev.Monsters {
			if m.Position == newPoint {
				return false
			}
		}

		m.Position.move(dir)

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
