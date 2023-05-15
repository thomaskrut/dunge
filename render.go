package main

func render(d *dungeon, p player, monsters map[Point]*monster, items map[Point]*item) (toPrint []rune) {

	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {

			var char rune

			if i, ok := items[Point{x, y}]; ok && d.grid[x][y]&lit == lit {
				char = i.getChar()
			}

			if m, ok := monsters[Point{x, y}]; ok && d.grid[x][y]&lit == lit {
				char = m.getChar()
			}

			if char == 0 {
				char = charmap.chars[(d.getPoint(Point{x, y}))]
			}
			
			if p.getPosition().overlaps(Point{x, y}) {
				char = p.getChar()
			}
			toPrint = append(toPrint, char)
		}
		toPrint = append(toPrint, '\n')
	}

	return toPrint

}
