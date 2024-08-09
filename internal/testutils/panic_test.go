package testutils_test

import (
	"errors"
	"testing"

	"github.com/mpstewart/go-helpers/internal/testutils"
)

func TestDidPanic(t *testing.T) {
	t.Run("panics", func(t *testing.T) {
		wantErr := errors.New("this is an error")
		did, gotErr := testutils.DidPanic(func() {
			panic(wantErr)
		})

		// compilation assertion that DidPanic returns an any type error,
		// because that's what recover() returns
		var _ any = gotErr

		if !did {
			t.Errorf("failed to detect a panic")
		}

		if gotErr == nil {
			t.Fatalf("failed to return the provided panic message")
		}

		if gotErr := gotErr.(error); !errors.Is(gotErr, wantErr) {
			t.Errorf("got err=%s, expected err err=%s", gotErr, wantErr)
		}
	})

	t.Run("no panic", func(t *testing.T) {
		did, gotErr := testutils.DidPanic(func() {
		})
		if did {
			t.Error("detected an unexpected panic")
		}

		if gotErr != nil {
			t.Errorf("returned an unexpected error: %v", gotErr)
		}
	})
}
