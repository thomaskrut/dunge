package main

type arrowQueue struct {
	arrows []point
}

func (a *arrowQueue) push(p point) {
	a.arrows = append(a.arrows, p)
}

func (a *arrowQueue) pop() point {
	if len(a.arrows) == 0 {
		return point{-1, -1}
	}
	p := a.arrows[0]
	a.arrows = a.arrows[1:]
	return p
}