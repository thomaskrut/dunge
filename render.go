package main

func render(d *dungeonMap, p player, arrows *arrowQueue, overlay []string, viewportWidth, viewportHeight int, monsters map[point]*monster, items map[point][]*item, features map[point]*feature) (toPrint []rune) {

	viewportHeight /= 2
	viewportWidth /= 2

	rowCounter := 0
	charCounter := 0

	arrow := arrows.pop()

	for y := p.position.y - viewportHeight; y < p.position.y+viewportHeight; y++ {

		for x := p.position.x - viewportWidth; x < p.position.x+viewportWidth; x++ {

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

			if f, ok := features[point{x, y}]; ok && (d.grid[x][y]&visited == visited || d.grid[x][y]&lit == lit) {
				char = f.getChar()
			}

			if i, ok := items[point{x, y}]; ok && d.grid[x][y]&lit == lit {
				char = i[0].getChar()
			}

			if m, ok := monsters[point{x, y}]; ok && d.grid[x][y]&lit == lit {
				char = m.getChar()
			}

			if char == 0 {
				char = charmap.chars[(d.read(point{x, y}))]
			}

			if p.getPosition() == (point{x, y}) {
				char = p.getChar()
			}

			if arrow.x == x && arrow.y == y {
				char = '^'
			}

			toPrint = append(toPrint, char)
		}
		toPrint = append(toPrint, '\n')
		rowCounter++
		charCounter = 0
	}

	return toPrint

}

func renderAll(d *dungeonMap, p player, arrows *arrowQueue, overlay []string, monsters map[point]*monster, items map[point]*item, features map[point]*feature) (toPrint []rune) {

	rowCounter := 0
	charCounter := 0

	arrow := arrows.pop()

	for y := 0; y < d.height; y++ {

		for x := 0; x < d.width; x++ {

			d.grid[x][y] = d.grid[x][y] | visited

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

			if f, ok := features[point{x, y}]; ok && (d.grid[x][y]&visited == visited || d.grid[x][y]&lit == lit) {
				char = f.getChar()
			}

			if i, ok := items[point{x, y}]; ok && d.grid[x][y]&lit == lit {
				char = i.getChar()
			}

			if m, ok := monsters[point{x, y}]; ok && d.grid[x][y]&lit == lit {
				char = m.getChar()
			}

			if char == 0 {
				char = charmap.chars[(d.read(point{x, y}))]
			}

			if p.getPosition() == (point{x, y}) {
				char = p.getChar()
			}

			if arrow.x == x && arrow.y == y {
				char = '^'
			}

			toPrint = append(toPrint, char)
		}
		toPrint = append(toPrint, '\n')
		rowCounter++
		charCounter = 0
	}

	return toPrint

}
