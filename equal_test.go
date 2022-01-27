package floattest_test

import (
	"testing"

	"github.com/crhntr/floattest"
)

func TestEqual(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		isMatch, err := floattest.Equal(3, 9.54321).Match(9.543001)
		if err != nil {
			t.Error("unexpected error", err)
		}
		if !isMatch {
			t.Error("expected match")
		}
	})
	t.Run("not match", func(t *testing.T) {
		isMatch, err := floattest.Equal(3, 9.54321).Match(9.5)
		if err != nil {
			t.Error("unexpected error", err)
		}
		if isMatch {
			t.Error("expected not match")
		}
	})
	t.Run("not match rounded", func(t *testing.T) {
		isMatch, err := floattest.Equal(3, 9.54321).Match(9.54399)
		if err != nil {
			t.Error("unexpected error", err)
		}
		if isMatch {
			t.Error("expected not match")
		}
	})
	t.Run("not float", func(t *testing.T) {
		_, err := floattest.Equal(3, 9.54321).Match(struct{}{})
		if err == nil {
			t.Error("expected error", err)
		}
	})
	t.Run("failure message", func(t *testing.T) {
		floattest.Equal(3, 9.54321).FailureMessage(9.54399)
	})
	t.Run("negated failure message", func(t *testing.T) {
		floattest.Equal(3, 9.54321).NegatedFailureMessage(9.54399)
	})
}

func TestEqualSlice(t *testing.T) {
	t.Run("match", func(t *testing.T) {
		isMatch, err := floattest.EqualSlice(3, []float64{9.54321}).Match([]float64{9.543001})
		if err != nil {
			t.Error("unexpected error", err)
		}
		if !isMatch {
			t.Error("expected match")
		}
	})
	t.Run("not match", func(t *testing.T) {
		isMatch, err := floattest.EqualSlice(3, []float64{9.54321}).Match([]float64{9.5})
		if err != nil {
			t.Error("unexpected error", err)
		}
		if isMatch {
			t.Error("expected not match")
		}
	})
	t.Run("not match rounded", func(t *testing.T) {
		isMatch, err := floattest.EqualSlice(3, []float64{9.54321}).Match([]float64{9.54399})
		if err != nil {
			t.Error("unexpected error", err)
		}
		if isMatch {
			t.Error("expected not match")
		}
	})
	t.Run("wrong len", func(t *testing.T) {
		isMatch, err := floattest.EqualSlice(3, []float64{9.54321}).Match([]float64(nil))
		if err != nil {
			t.Error("unexpected error", err)
		}
		if isMatch {
			t.Error("expected not match")
		}
	})
	t.Run("not float", func(t *testing.T) {
		_, err := floattest.EqualSlice(3, []float64{9.54321}).Match(struct{}{})
		if err == nil {
			t.Error("expected error", err)
		}
	})
	t.Run("failure message", func(t *testing.T) {
		floattest.EqualSlice(3, []float64{9.54321}).FailureMessage([]float64{9.54399})
	})
	t.Run("negated failure message", func(t *testing.T) {
		floattest.EqualSlice(3, []float64{9.54321}).NegatedFailureMessage([]float64{9.54399})
	})
	t.Run("failure message", func(t *testing.T) {
		floattest.EqualSlice(3, []float64{9.54321}).FailureMessage([]float64(nil))
	})
	t.Run("negated failure message", func(t *testing.T) {
		floattest.EqualSlice(3, []float64{9.54321}).NegatedFailureMessage([]float64(nil))
	})
}
