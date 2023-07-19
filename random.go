package main

import (	
	"time"
	"math/rand"
)

var (
	source rand.Source
)

func setSource(seed int) {
	if seed == 0 {
		source = rand.NewSource(time.Now().UnixNano())
	} else {
		source = rand.NewSource(int64(seed))
	}
}

func randomNumber(size int) int {
	return rand.New(source).Intn(size)
}