package objects

type Size struct {
	W int32
	H int32
}

type Position struct {
	X int32
	Y int32
}

type Color struct {
	Red     uint8
	Green   uint8
	Blue    uint8
	Opacity uint8
}

func (s *Size) SetSize(w int32, h int32) {
	s.W = w
	s.H = h
}

func (p *Position) SetPosition(x int32, y int32) {
	p.X = x
	p.Y = y
}

func (c *Color) SetColor(red uint8, green uint8, blue uint8, opacity uint8) {
	c.Red = red
	c.Green = green
	c.Blue = blue
	c.Opacity = opacity
}
