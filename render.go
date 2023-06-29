package main

func render(d *dungeon, p player, overlay []string, viewportWidth, viewportHeight int, monsters map[point]*monster, items map[point]*item, features map[point]*feature) (toPrint []rune) {

	viewportHeight /= 2
	viewportWidth /= 2

	rowCounter := 0
	charCounter := 0

	for y := p.position.y - viewportHeight; y < p.position.y + viewportHeight; y++ {

		for x := p.position.x - viewportWidth; x < p.position.x + viewportWidth; x++ {

			if len(overlay) > rowCounter {
				if len(overlay[rowCounter]) > charCounter {
					toPrint = append(toPrint, rune(overlay[rowCounter][charCounter]))
					charCounter++
					continue
				}
			}

			if x < 0 || x >= d.width || y < 0 || y >= d.height {
				toPrint = append(toPrint, ' ')
				continue
			}

			var char rune

			if f, ok := features[point{x, y}]; ok && (d.grid[x][y]&visited == visited || d.grid[x][y]&lit == lit) && d.grid[x][y]&arrow != arrow {
				char = f.getChar()
			}

			if i, ok := items[point{x, y}]; ok && d.grid[x][y]&lit == lit && d.grid[x][y]&arrow != arrow {
				char = i.getChar()
			}

			if m, ok := monsters[point{x, y}]; ok && d.grid[x][y]&lit == lit && d.grid[x][y]&arrow != arrow {
				char = m.getChar()
			}

			if char == 0 {
				char = charmap.chars[(d.getPoint(point{x, y}))]
			}

			if p.getPosition() == (point{x, y}) && d.grid[x][y]&arrow != arrow {
				char = p.getChar()
			}
			toPrint = append(toPrint, char)
		}
		toPrint = append(toPrint, '\n')
		rowCounter++
		charCounter = 0
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
