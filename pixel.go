package newyeartree

func NewPixel(x, y int, c Color) (p Pixel) {
	p.X = x
	p.Y = y
	p.Color = c
	p.Brightness = 255
	return
}

type Pixel struct {
	X, Y       int
	Brightness uint32
	Color
}

func (p Pixel) Clone() (r Pixel) {
	r.X, r.Y = p.X, p.Y
	r.Color = p.Color
	return
}

func (p Pixel) Get(x, y int) Color {
	if x == p.X && y == p.Y {
		p.Red = p.Red * p.Brightness >> 8
		p.Green = p.Green * p.Brightness >> 8
		p.Blue = p.Blue * p.Brightness >> 8
		return p.Color
	}
	return Color{}
}
