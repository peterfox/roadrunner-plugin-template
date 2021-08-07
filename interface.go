package plugin

// Actionable is an interface that can be used in additional plugins via endure container
type Actionable interface {
	// Action is purely an example of an interface available to Endure
	Action() error
}
