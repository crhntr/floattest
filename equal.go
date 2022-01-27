package floattest

import (
	"fmt"
	"math"
	"strings"
)

func roundFloat64(precision int, num float64) float64 {
	shift := math.Pow(10, float64(precision))
	shifted := num * shift
	rounded := int(shifted + math.Copysign(0.5, shifted))
	return float64(rounded) / shift
}

func Equal(precision int, expect float64) equalFloat {
	return equalFloat{
		places: precision,
		exp:    expect,
	}
}

type equalFloat struct {
	places int
	exp    float64
}

func (ef equalFloat) Match(actual interface{}) (success bool, err error) {
	a, ok := actual.(float64)
	if !ok {
		return false, fmt.Errorf("expected type %T but got %T", ef.exp, actual)
	}

	exp := roundFloat64(ef.places, ef.exp)
	got := roundFloat64(ef.places, a)

	return exp == got, nil
}

func (ef equalFloat) FailureMessage(actual interface{}) (message string) {
	a := actual.(float64)

	exp := roundFloat64(ef.places, ef.exp)
	got := roundFloat64(ef.places, a)

	return fmt.Sprintf("Expected %f to equal %f", got, exp)
}

func (ef equalFloat) NegatedFailureMessage(actual interface{}) (message string) {
	a := actual.(float64)

	exp := roundFloat64(ef.places, ef.exp)
	got := roundFloat64(ef.places, a)

	return fmt.Sprintf("Expected %f not to equal %f", got, exp)
}

func EqualSlice(precision int, expect []float64) equalFloats {
	return equalFloats{
		places: precision,
		exp:    expect,
	}
}

type equalFloats struct {
	places int
	exp    []float64
}

func (ef equalFloats) Match(actual interface{}) (success bool, err error) {
	a, ok := actual.([]float64)
	if !ok {
		return false, fmt.Errorf("expected type %T but got %T", ef.exp, actual)
	}

	if len(a) != len(ef.exp) {
		return false, nil
	}

	for i := range ef.exp {
		exp := roundFloat64(ef.places, ef.exp[i])
		got := roundFloat64(ef.places, a[i])
		if exp != got {
			return false, nil
		}
	}

	return true, nil
}

func (ef equalFloats) FailureMessage(actual interface{}) (message string) {
	a := actual.([]float64)

	if len(a) != len(ef.exp) {
		return fmt.Sprintf("Expected len %v to have len %d but got %d", a, len(ef.exp), len(a))
	}

	b := strings.Builder{}
	for i := range ef.exp {
		exp := roundFloat64(ef.places, ef.exp[i])
		got := roundFloat64(ef.places, a[i])
		if exp != got {
			b.WriteString(fmt.Sprintf("Expected value at index %d to equal %f but got %f\n", i, exp, got))
		}
	}

	return b.String()
}

func (ef equalFloats) NegatedFailureMessage(actual interface{}) (message string) {
	a := actual.([]float64)

	if len(a) != len(ef.exp) {
		return fmt.Sprintf("Expected len %v not to have len %d but got %d", a, len(ef.exp), len(a))
	}

	b := strings.Builder{}
	for i := range ef.exp {
		exp := roundFloat64(ef.places, ef.exp[i])
		got := roundFloat64(ef.places, a[i])
		if exp != got {
			b.WriteString(fmt.Sprintf("Expected value at index %d to not equal %f but got %f\n", i, exp, got))
		}
	}

	return b.String()
}
