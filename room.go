package main

type scannedRoom struct {
	points []point
}

func (r *scannedRoom) add(p point) {
	r.points = append(r.points, p)
}

func (r *scannedRoom) clear() {
	r.points = nil
}
