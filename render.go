package main

func render(d *dungeon, p player, monsters []monster) (toPrint []rune) {

	
	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			var char rune
			for _, m := range monsters {
				if p.getPosition().overlaps(Point{x, y}) {
					char = p.getChar()
					break
				} else if m.getPosition().overlaps(Point{x, y}) {
					char = m.getChar()
					break
				}
			}
			if char == 0 {
				char = charmap.chars[(d.getPoint(Point{x, y}))]
			}
			toPrint = append(toPrint, char)
		}
		toPrint = append(toPrint, '\n')
	}

	return toPrint

}
