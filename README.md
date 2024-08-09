# go-helpers

Misc utilities for the Go programming language. Don't use these.

These are useful when writing table tests where literals may be difficult to
initialize because the init functions either or return incorrect types, or
cannot throw errors the way you are calling them (but still want to find out if
they start).

Contrived example in `example/examples_test.go`.