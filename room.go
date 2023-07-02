package main

type currentRoom struct {
	points []point
}

func (r *currentRoom) add(p point) {
	r.points = append(r.points, p)
}

func (r *currentRoom) clear() {
	r.points = nil
}