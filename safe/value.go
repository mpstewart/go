// safe provides helper functions for safely checking pointers
package safe

// Value takes a *T, and tests if for nil. If it is nil, a zero value of T is
// returned. If it is not nil, it is de-referenced and the underlying value is
// returned
func Value[T any](t *T) T {
	if t == nil {
		var zero T
		return zero
	}

	return *t
}
