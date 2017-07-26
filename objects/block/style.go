package block

const (
	// Filled style
	Filled = 1
	// Border to draw only border
	Border = 2
)

// Styler define personalisable style of block
type Styler struct {
	block uint8
}
