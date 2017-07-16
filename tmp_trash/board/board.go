package board

import "time"

const (
	ValueEmpty = 0
	ValueP1    = 1
	ValueP2    = 2
)

type Board struct {
	Grid [][]Cell
}

type Cell struct {
	Value  uint8
	Weight uint8
}

type Move struct {
	X      uint8
	Y      uint8
	Reflex time.Duration // time to play
}
