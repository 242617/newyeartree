package newyeartree

type Color struct {
	Red   uint32
	Green uint32
	Blue  uint32
}

func (c *Color) SetBrightness(a uint32) {
	c.Red = c.Red * a >> 8
	c.Green = c.Green * a >> 8
	c.Blue = c.Blue * a >> 8
}

func (c *Color) Mix(w Color) {
	c.Red += w.Red
	c.Green += w.Green
	c.Blue += w.Blue
	c.normalize()
}

func (c Color) Equals(t Color) bool {
	return (t.Red == c.Red) && (t.Green == c.Green) && (t.Blue == c.Blue)
}

func (c Color) Clone() (r Color) {
	r.Red = c.Red
	r.Green = c.Green
	r.Blue = c.Blue
	return
}

func (c *Color) normalize() {
	if c.Red >= 255 {
		c.Red = 255
	}
	if c.Green >= 255 {
		c.Green = 255
	}
	if c.Blue >= 255 {
		c.Blue = 255
	}
}
