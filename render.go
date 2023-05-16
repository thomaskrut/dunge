package main

func render(d *dungeon, p player, monsters map[point]*monster, items map[point]*item) (toPrint []rune) {

	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {

			var char rune

			if i, ok := items[point{x, y}]; ok && d.grid[x][y]&lit == lit {
				char = i.getChar()
			}

			if m, ok := monsters[point{x, y}]; ok && d.grid[x][y]&lit == lit {
				char = m.getChar()
			}

			if char == 0 {
				char = charmap.chars[(d.getPoint(point{x, y}))]
			}

			if p.getPosition() == (point{x, y}) {
				char = p.getChar()
			}
			toPrint = append(toPrint, char)
		}
		toPrint = append(toPrint, '\n')
	}

	return toPrint

}

func renderAll(d *dungeon, p player, monsters map[point]*monster, items map[point]*item) (toPrint []rune) {

	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			d.grid[x][y] = d.grid[x][y] | visited
			var char rune

			if i, ok := items[point{x, y}]; ok {
				char = i.getChar()
			}

			if m, ok := monsters[point{x, y}]; ok {
				char = m.getChar()
			}

			if char == 0 {
				char = charmap.chars[d.getPoint(point{x, y})]
			}

			if p.getPosition() == (point{x, y}) {
				char = p.getChar()
			}
			toPrint = append(toPrint, char)
		}
		toPrint = append(toPrint, '\n')
	}

	return toPrint

}
