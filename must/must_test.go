package must_test

import (
	"errors"
	"testing"

	"github.com/mpstewart/go-helpers/internal/testutils"
	"github.com/mpstewart/go-helpers/must"
)

const (
	no  = false
	yes = true
	n   = 999
)

var err = errors.New("was told to error")

func maybeErrors(shouldError bool) error {
	if shouldError {
		return err
	}

	return nil
}

func maybeReturns(i int, err error) (int, error) {
	if err != nil {
		return 0, err
	}

	return i + n, nil
}

func TestLive(t *testing.T) {
	if panicked, _ := testutils.DidPanic(func() {
		must.Live(maybeErrors(no))
	}); panicked {
		t.Errorf("panicked on a non-error")
	}

	if did, msg := testutils.DidPanic(func() {
		must.Live(maybeErrors(yes))
	}); !did {
		t.Errorf("failed to panic on an error")
		if !errors.Is(msg.(error), err) {
			t.Errorf("got=%s, expected=%s", msg, err)
		}
	}
}

func TestReturn(t *testing.T) {
	if panicked, _ := testutils.DidPanic(func() {
		_ = must.Return(maybeReturns(15, maybeErrors(yes)))
	}); !panicked {
		t.Fatal("did not panic on error")
	}

	if panicked, _ := testutils.DidPanic(func() {
		_ = must.Return(maybeReturns(15, maybeErrors(no)))
	}); panicked {
		t.Fatal("panicked when no error")
	}
}
