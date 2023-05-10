package main

type entity interface {
	
	getPosition() Point
	
	getChar() rune

	move(direction) bool
	
	moveCounter() float32
}
