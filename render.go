package main


func render(d *dungeon, p player, monsters map[Point]monster, items []item) (toPrint []rune) {

	
	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {

			var char rune

			if m, ok := monsters[Point{x, y}]; ok {
					char = m.getChar()
				}

			for _, i := range items {
				if i.getPosition().overlaps(Point{x, y}) {
					char = i.getChar()
					break
				}
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
