package main

type characterMapper struct {
	chars map[byte]rune
}

func initChapMap() characterMapper {
	charmap := characterMapper{make(map[byte]rune)}
	charmap.add(obstacle, ' ')
	charmap.add(empty, ' ')
	charmap.add(empty|room, ' ')
	charmap.add(empty|room|visited, ' ')
	charmap.add(empty|visited, ' ')
	charmap.add(empty|lit, '.')
	charmap.add(empty|room|lit, '.')
	charmap.add(visited|obstacle, 9617)
	charmap.add(lit|obstacle, 9618)

	return charmap
}

func (c characterMapper) add(value byte, char rune) {
	c.chars[value] = char
}
