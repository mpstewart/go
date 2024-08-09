package safe_test

import (
	"testing"

	"github.com/mpstewart/go-helpers/safe"
)

var (
	i    = 1
	iptr = &i
)

type foo struct {
	value int
}

func TestValue(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("safe.Value does not panic on non-nil")
		}
	}()

	var p *int
	f := foo{
		value: safe.Value(p),
	}
	if got := f.value; got != 0 {
		t.Errorf("got=%d, expected=%d", 0, got)
	}

	f = foo{
		value: safe.Value(iptr),
	}
	if got := f.value; got != *iptr {
		t.Errorf("got=%d, expected=%d", got, *iptr)
	}
}
