package widget

// Alpha 0: alpha disable
// Alpha 0: min oppacity
// Alpha 255: max oppacity
var (
	ColorWhite = Color{
		Red:   255,
		Green: 255,
		Blue:  255,
		Alpha: 255,
	}
	ColorBlack Color = Color{
		Red:   0,
		Green: 0,
		Blue:  0,
		Alpha: 255,
	}
	ColorRed Color = Color{
		Red:   255,
		Green: 0,
		Blue:  0,
		Alpha: 255,
	}
	ColorGreen Color = Color{
		Red:   0,
		Green: 255,
		Blue:  0,
		Alpha: 255,
	}
	ColorBlue Color = Color{
		Red:   0,
		Green: 0,
		Blue:  255,
		Alpha: 255,
	}
)

type Colors struct {
	base   *Color
	hover  *Color
	action *Color
}

type Color struct {
	Red   uint8
	Green uint8
	Blue  uint8
	Alpha uint8
}
