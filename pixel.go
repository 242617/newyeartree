package newyeartree

type Pixel struct {
	X, Y int
	Color
}

func (p Pixel) Clone() (r Pixel) {
	r.X, r.Y = p.X, p.Y
	r.Color = p.Color
	return
}

func (p Pixel) Get(x, y int) Color {
	if x == p.X && y == p.Y {
		return p.Color
	}
	return Color{}
}
