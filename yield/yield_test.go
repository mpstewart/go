package yield_test

import (
	"testing"

	"github.com/mpstewart/go-helpers/yield"
)

var (
	i    = 12
	iptr = &i
)

type foo struct {
	pointer *int
}

func TestPtr(t *testing.T) {
	f := foo{
		pointer: yield.Ptr(i),
	}

	if f.pointer == nil {
		t.FailNow()
	}

	if *f.pointer != i {
		t.Fail()
	}
}

func TestValue(t *testing.T) {
	t.Run("panics", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("yield.Value panics on nil")
			}
		}()

		_ = yield.Value[int](nil)
	})

	t.Run("no panic", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("yield.Value does not panic on non-nil")
			}
		}()

		gotVal := yield.Value(iptr)

		if gotVal != i {
			t.Errorf("got=%d, expected=%d", gotVal, i)
		}

	})
}
