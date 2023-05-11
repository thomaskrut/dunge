package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

type monsterList struct {
	Monsters []monster `json:"monsters"`
}

type monster struct {
	position         Point
	Char             string   `json:"char"`
	Name             string   `json:"name"`
	Prob             int      `json:"prob"`
	Attack           []string `json:"attack"`
	Str              int      `json:"str"`
	Hp               int      `json:"hp"`
	Moves            bool     `json:"moves"`
	Aggressive       bool     `json:"aggressive"`
	Speed            float32  `json:"speed"`
	moveCounterValue float32
}

func (m *monster) takeDamage(damage int) {
	m.Hp -= damage
	fmt.Println(m.Name, "took", damage, "damage and has", m.Hp, "hp left")
}

func (m *monster) attack(p *player) {
	if (m.Hp > 0) {
		messages.addMessage(m.Name + " " + m.Attack[rand.Intn(len(m.Attack))] + " you")
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

func (m monster) getPosition() Point {
	return m.position
}

func (m *monster) setPosition(point Point) {
	m.position = point
}

func (m monster) getChar() rune {
	return rune(m.Char[0])
}

func (m *monster) move(dir direction) bool {
	if m.position.getPossibleDirections(&d)[dir] {
		
		newPoint := m.position
		newPoint.new(dir)
		if (newPoint.overlaps(p.position)) {
			m.attack(&p)
			return true
		}

		m.position.new(dir)
		
		return true
	}
	return false
}

func readMonsterTemplate() monsterList {

	data, err := os.ReadFile("templates.json")
	if err != nil {
		panic(err)
	}

	fileAsString := string(data)

	monsterTemplates := monsterList{}

	if err := json.Unmarshal([]byte(fileAsString), &monsterTemplates); err != nil {
		panic(err)
	}

	fmt.Println(monsterTemplates)

	return monsterTemplates

}
