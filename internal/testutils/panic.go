package testutils

func DidPanic(fn func()) (did bool, err any) {
	defer func() {
		r := recover()
		did = r != nil
		err = r
	}()

	fn()

	return
}
