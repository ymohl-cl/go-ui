package data

// Model is interface to provide datas on your scenes
type Model interface {
	// Build
	Build() error
}
