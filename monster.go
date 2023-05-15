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

	position         point

	Char             string   `json:"char"`
	Name             string   `json:"name"`
	Prob             int      `json:"prob"`
	AttackVerbs      []string `json:"attack"`
	Str              int      `json:"str"`
	Hp               int      `json:"hp"`
	Moves            bool     `json:"moves"`
	Aggressive       bool     `json:"aggressive"`
	Speed            float32  `json:"speed"`
	Movesdiagonally  bool     `json:"movesdiagonally"`

	moveCounterValue float32

}

func (m *monster) takeDamage(damage int) {
	m.Hp -= damage
}

func (m *monster) attack(p *player) {
	if m.Hp > 0 {
		messages.push(m.Name + " " + m.AttackVerbs[randomNumber(len(m.AttackVerbs))] + " you")
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
		newPoint.new(dir)
		if newPoint == p.position {
			m.attack(&p)
			return true
		}
		for _, m := range activeMonsters {
			if m.position == newPoint {
				return false
			}
		}

		m.position.new(dir)

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
