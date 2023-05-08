package main

func render(d *dungeon, entities ...entity) (toPrint []rune) {

	for y := 0; y < d.height; y++ {
		for x := 0; x < d.width; x++ {
			var char rune
			for _, entity := range entities {
				if entity.getPosition().overlaps(Point{x, y}) {
					char = entity.getChar()
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
