package main

type characterMapper struct {
	chars map[int]rune
}

func newCharMap() characterMapper {
	return characterMapper{make(map[int]rune)}
}

func (c characterMapper) add(value int, char rune) {
	c.chars[value] = char
}
