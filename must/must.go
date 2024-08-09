// must provides helper functions that cause error-returning functions with
// common signatures to panic
package must

// Return takes an instance of T and an error. If the err is not nil, Return
// panics. If the error is not nil, the T value is returned.
func Return[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

// Live takes an error, and if it is nil, panics with the error message
func Live(err error) {
	if err != nil {
		panic(err)
	}
}
