// a collection of interface specifications for objects that are part of factor package
package factor

// the Factor interface defines actions that are accessible by the factor objects
type Factor interface {
	Calculate() (float64,error)
}