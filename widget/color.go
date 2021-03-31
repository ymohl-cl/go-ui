package widget

// Alpha 0: alpha disable
// Alpha 1: min oppacity
// Alpha 255: max oppacity
var (
	ColorWhite = Color{
		Red:   255,
		Green: 255,
		Blue:  255,
		Alpha: 0,
	}
	ColorBlack Color = Color{
		Red:   0,
		Green: 0,
		Blue:  0,
		Alpha: 0,
	}
	ColorRed Color = Color{
		Red:   255,
		Green: 0,
		Blue:  0,
		Alpha: 0,
	}
	ColorGreen Color = Color{
		Red:   0,
		Green: 255,
		Blue:  0,
		Alpha: 0,
	}
	ColorBlue Color = Color{
		Red:   0,
		Green: 0,
		Blue:  255,
		Alpha: 0,
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
