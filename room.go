package main

type scannedRoom struct {
	Points []point
}

func (r *scannedRoom) add(p point) {
	r.Points = append(r.Points, p)
}

func (r *scannedRoom) clear() {
	r.Points = nil
}
