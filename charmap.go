package main

type characterMapper struct {
	chars map[int]rune
}

func NewCharMap() characterMapper {
	return characterMapper{make(map[int]rune)}
}

func (c characterMapper) Add(value int, char rune) {
	c.chars[value] = char
}
