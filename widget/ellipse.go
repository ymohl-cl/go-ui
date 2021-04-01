package widget

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	EllipseStyleFill EllipseStyle = iota
	EllipseStyleBorder
)

type EllipseStyle uint8

type Ellipse struct {
	widget
	style  EllipseStyle
	points []Position
}

func NewEllipse(w, h int32) *Ellipse {
	e := Ellipse{}

	e.SetSize(w, h)
	e.setPoints()
	return &e
}

func (e *Ellipse) Close() {}

func (e *Ellipse) SetStyle(s EllipseStyle) {
	e.style = s
}

// algo to check point is inside ellipse
// https://www.geeksforgeeks.org/check-if-a-point-is-inside-outside-or-on-the-ellipse/
func (e Ellipse) isInside(x, y int32) bool {
	v := math.Pow(float64(e.position.X)-float64(x), 2.0)/math.Pow(float64(e.block.width)/2.0, 2.0) + (math.Pow((float64(y)-float64(e.position.Y)), 2.0) / math.Pow(float64(e.block.height)/2.0, 2.0))
	if v <= 1.0 {
		return true
	}
	return false
}

// IsHover override the widget IsHover to check is mouse is hover the circle
func (e Ellipse) IsHover(x, y int32) bool {
	return e.isInside(x, y)
}

func (e *Ellipse) SetSize(w, h int32) {
	e.widget.SetSize(w, h)
	e.points = []Position{}
	e.setPoints()
}

func (e *Ellipse) SetPosition(x, y int32) {
	e.widget.SetPosition(x, y)
	e.points = []Position{}
	e.setPoints()
}

func (e *Ellipse) renderFill(r *sdl.Renderer) error {
	midX := e.block.width / 2
	maxX := e.position.X + midX
	midY := e.block.height / 2
	maxY := e.position.Y + midY
	for x := e.position.X - midX; x <= maxX; x++ {
		for y := e.position.Y - midY; y <= maxY; y++ {
			if !e.isInside(x, y) {
				continue
			}
			if err := r.DrawPoint(x, y); err != nil {
				return err
			}
		}
	}
	return nil
}

func (e *Ellipse) renderBorder(r *sdl.Renderer) error {
	var err error

	x := e.position.X
	y := e.position.Y
	for _, p := range e.points {
		if err = r.DrawPoint(p.X, p.Y); err != nil {
			return err
		}
		xp := x - (p.X - x)
		yp := y - (p.Y - y)
		if err = r.DrawPoint(xp, yp); err != nil {
			return err
		}
	}
	return nil
}

func (e *Ellipse) Render(r *sdl.Renderer) error {
	var err error

	c := e.Color(e.state)
	if err = r.SetDrawColor(c.Red, c.Green, c.Blue, c.Alpha); err != nil {
		return err
	}

	if e.style == EllipseStyleFill {
		if err = e.renderFill(r); err != nil {
			return err
		}
		return nil
	}
	// else
	if err = e.renderBorder(r); err != nil {
		return err
	}
	return nil
}

// setPoints ellipse to draw wit Midpoint ellipse drawing
// https://www.geeksforgeeks.org/midpoint-ellipse-drawing-algorithm/
func (e *Ellipse) setPoints() {
	//	var err error
	rx := float32(e.block.width) / 2
	ry := float32(e.block.height) / 2

	var dx, dy, d1, d2, x, y float32
	y = float32(ry)

	// Initial decision parameter of region 1
	d1 = (ry * ry) - (rx * rx * ry) + (0.25 * rx * rx)
	dx = 2 * ry * ry * x
	dy = 2 * rx * rx * y

	// For region 1
	for dx < dy {
		e.points = append(e.points, Position{X: int32(x) + e.position.X, Y: int32(y) + e.position.Y})
		e.points = append(e.points, Position{X: int32(x) + e.position.X, Y: int32(-y) + e.position.Y})

		// Checking and updating value of
		// Decision parameter bases on algorithm
		x++
		if d1 < 0 {
			dx = dx + (2.0 * ry * ry)
			d1 = d1 + dx + (ry * ry)
		} else {
			y--
			dx = dx + (2.0 * ry * ry)
			dy = dy - (2.0 * rx * rx)
			d1 = d1 + dx - dy + (ry * ry)
		}
	}

	// Decision parameter of region 2
	d2 = ((ry * ry) * ((x + 0.5) * (x + 0.5))) + ((rx * rx) * ((y - 1) * (y - 1))) - (rx * rx * ry * ry)

	// For region 2
	for y >= 0 {
		// print points bases on 4-way symmetry
		e.points = append(e.points, Position{X: int32(x) + e.position.X, Y: int32(y) + e.position.Y})
		e.points = append(e.points, Position{X: int32(x) + e.position.X, Y: int32(-y) + e.position.Y})

		// Checking and updating parameter
		// value based on algorithm
		y--
		if d2 > 0 {
			dy = dy - (2.0 * rx * rx)
			d2 = d2 + (rx * rx) - dy
		} else {
			x++
			dx = dx + (2.0 * ry * ry)
			dy = dy - (2.0 * rx * rx)
			d2 = d2 + dx - dy + (rx * rx)
		}
	}
}
