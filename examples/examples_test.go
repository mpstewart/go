package main

import (
	"math/rand/v2"
	"strconv"
	"testing"

	"github.com/mpstewart/go-helpers/must"
	"github.com/mpstewart/go-helpers/safe"
	"github.com/mpstewart/go-helpers/yield"
)

// Foo is a contrivance of a struct that is hard to initialize because of
// pointers and weird struct literals
type Foo struct {
	Bar *int
}

func (f Foo) BarIsEven() bool {
	if f.Bar == nil {
		return false
	}
	return *f.Bar%2 == 0

}

func TestFooBarIsEven(t *testing.T) {
	for _, f := range []Foo{
		// it is super annoying to initialize these literally

		{ // in exactly this one case, it's not too bad
			Bar: nil,
		},

		{ // this is possible, but ugly
			Bar: func() *int {
				i := 15
				return &i
			}(),
		},

		{ // much nicer
			Bar: yield.Ptr(15),
		},

		{ // it even works with function calls
			Bar: yield.Ptr(rand.Int()),
		},
		{ // even function calls that are error-prone
			Bar: yield.Ptr(must.Return(strconv.Atoi("15"))),
		},
		// { // panics!
		// 	Bar: yield.Ptr(must.Return(strconv.Atoi("foo"))),
		// },
		{ // But this doesn't
			Bar: yield.Ptr(safe.Value[int](nil)),
		},
	} {
		if f.Bar == nil {
			continue
		}
		if (*f.Bar%2 == 0) != f.BarIsEven() {
			t.Errorf("BarIsEven busted for %d", *f.Bar)

		}
	}

}
