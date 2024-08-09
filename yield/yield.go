// yield provides helper functions to use when writing tests that build literal
// structs where values are initialized with function calls that return the
// wrong pointer/value types.
package yield

// Ptr will, given a value of type T, return a pointer to the value
func Ptr[T any](t T) *T {
	return &t
}

// Value will, given a *T, de-reference it and return the underlying value. Will
// panic if *T is nil. See safe.Value if you want a zero value instead of a
// panic.
func Value[T any](t *T) T {
	return *t
}
