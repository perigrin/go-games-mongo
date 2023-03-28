package main

type Rect struct {
	X1 int
	X2 int
	Y1 int
	Y2 int
}

func NewRect(x, y, width, height int) Rect {
	return Rect{
		X1: x,
		Y1: y,
		X2: x + width,
		Y2: y + height,
	}
}

func (r *Rect) Center() (int, int) {
	x := (r.X1 + r.X2) / 2
	y := (r.Y1 + r.Y2) / 2
	return x, y
}

func (r *Rect) Intersect(o Rect) bool {
	if r.X2 < o.X1 || o.Y2 < r.X1 || r.X2 < o.Y1 || o.Y2 < r.Y1 {
		return false
	}
	return true
}
