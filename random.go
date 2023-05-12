package main

import (	
	"time"
	"math/rand"
)

var (
	source = rand.NewSource(time.Now().UnixNano())
)

func randomNumber(size int) int {
	return rand.New(source).Intn(size)
}