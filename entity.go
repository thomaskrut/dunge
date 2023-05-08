package main

type entity interface {
	move(dir direction)
	getPosition() Point
	setPosition(Point)
	getChar() rune
}
