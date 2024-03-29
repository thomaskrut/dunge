package main

func render(d *level, p player, arrows *arrowQueue, overlay []string, viewportWidth, viewportHeight int) (toPrint []rune) {

	viewportHeight /= 2
	viewportWidth /= 2

	rowCounter := 0
	charCounter := 0
	overlayMargin := 6

	arrow := arrows.pop()

	for y := p.Position.Y - viewportHeight; y < p.Position.Y+viewportHeight; y++ {

		for x := p.Position.X - viewportWidth; x < p.Position.X+viewportWidth; x++ {

			if rowCounter >= overlayMargin && len(overlay) > rowCounter-overlayMargin {
				if charCounter >= overlayMargin && len(overlay[rowCounter-overlayMargin]) > charCounter-overlayMargin {
					toPrint = append(toPrint, rune(overlay[rowCounter-overlayMargin][charCounter-overlayMargin]))
					charCounter++
					continue
				}
				charCounter++

			}

			if x < 0 || x >= d.Width || y < 0 || y >= d.Height {
				toPrint = append(toPrint, ' ')
				continue
			}

			var char rune

			if f, ok := lvl.Features[point{x, y}]; ok && (d.Grid[x][y]&visited == visited || d.Grid[x][y]&lit == lit) {
				char = f.getChar()
			}

			if i, ok := lvl.Items[point{x, y}]; ok && d.Grid[x][y]&lit == lit {
				char = i[len(lvl.Items[point{x, y}])-1].getChar()
			}

			if m, ok := lvl.Monsters[point{x, y}]; ok && d.Grid[x][y]&lit == lit {
				char = m.getChar()
			}

			if char == 0 {
				char = charmap.chars[(d.read(point{x, y}))]
			}

			if p.getPosition() == (point{x, y}) {
				char = p.getChar()
			}

			if arrow.X == x && arrow.Y == y {
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

func renderAll(d *level, p player, arrows *arrowQueue, monsters map[point]*monster, items map[point][]*item, features map[point]*feature) (toPrint []rune) {

	
	arrow := arrows.pop()

	for y := 0; y < d.Height; y++ {

		for x := 0; x < d.Width; x++ {

			var char rune

			if f, ok := features[point{x, y}]; ok && (d.Grid[x][y]&visited == visited || d.Grid[x][y]&lit == lit || d.Grid[x][y]&empty == empty) {
				char = f.getChar()
			}

			if i, ok := items[point{x, y}]; ok && d.Grid[x][y]&lit == lit {
				char = i[len(items[point{x, y}])-1].getChar()
			}

			if m, ok := monsters[point{x, y}]; ok && d.Grid[x][y]&lit == lit {
				char = m.getChar()
			}

			if char == 0 {
				char = charmap.chars[(d.read(point{x, y}))]
			}

			if p.getPosition() == (point{x, y}) {
				char = p.getChar()
			}

			if arrow.X == x && arrow.Y == y {
				char = '^'
			}

			toPrint = append(toPrint, char)
		}
		toPrint = append(toPrint, '\n')
	}

	return toPrint

}
