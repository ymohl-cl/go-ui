package data

// Model is interface to provide datas on your scenes
type Model interface {
	// Build the data model needest to the application
	Build() error
	// Close ressources used by the application
	Close() error
}
