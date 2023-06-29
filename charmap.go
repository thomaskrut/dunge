package main

type characterMapper struct {
	chars map[int]rune
}

func initChapMap() characterMapper {
	charmap := characterMapper{make(map[int]rune)}
	charmap.add(wall, ' ')
	charmap.add(empty, ' ')
	charmap.add(empty|visited, ' ')
	charmap.add(empty|lit, '.')
	charmap.add(visited|wall, 9617)
	charmap.add(lit|wall, 9618)
	charmap.add(lit|empty|arrow, '^')
	charmap.add(lit|wall|arrow, '^')
	charmap.add(visited|wall|arrow, '^')
	charmap.add(visited|empty|arrow, '^')
	
	/*
		charmap.add(visited|wall, 9637)
		charmap.add(lit|wall, 9639)
	*/
	return charmap
}

func (c characterMapper) add(value int, char rune) {
	c.chars[value] = char
}
