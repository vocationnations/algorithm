// defines interfaces that have to be satisfied by components
package components

// Component defines the operations that apply to the factor components.
type Component interface {
	// Calculate performs the calculation of a particular factor component
	CalculateFactor() (float64, error)
}
