package minmax_test

import (
	"math"
	"minmax"
	"testing"
)

func TestMinZeroElements(t *testing.T) {
	var list []int
	min, ok := minmax.Min(list...)

	if min != math.MaxInt {
		t.Errorf("Expected: %d, Got: %d", math.MaxInt, min)
	}
	if ok {
		t.Errorf("Expected: %t, Got: %t", false, ok)
	}
}

func TestMinNegativeElements(t *testing.T) {
	list := []int{-10, -20, -60, -50, -40}
	expected := -60
	min, ok := minmax.Min(list...)

	if min != expected {
		t.Errorf("Expected: %d, Got: %d", expected, min)
	}
	if !ok {
		t.Errorf("Expected: %t, Got: %t", true, ok)
	}
}

func TestMinPositiveElements(t *testing.T) {
	list := []int{10, 20, 60, 50, 40}
	expected := 10
	min, ok := minmax.Min(list...)

	if min != expected {
		t.Errorf("Expected: %d, Got: %d", expected, min)
	}
	if !ok {
		t.Errorf("Expected: %t, Got: %t", true, ok)
	}
}

func TestMinMixedElements(t *testing.T) {
	list := []int{80, -40, -50, 10, 30, 100}
	expected := -50
	min, ok := minmax.Min(list...)

	if min != expected {
		t.Errorf("Expected: %d, Got: %d", expected, min)
	}
	if !ok {
		t.Errorf("Expected: %t, Got: %t", true, ok)
	}
}

func TestMaxZeroElements(t *testing.T) {
	var list []int
	min, ok := minmax.Max(list...)

	if min != math.MinInt {
		t.Errorf("Expected: %d, Got: %d", math.MinInt, min)
	}
	if ok {
		t.Errorf("Expected: %t, Got: %t", false, ok)
	}
}

func TestMaxNegativeElements(t *testing.T) {
	list := []int{-10, -20, -60, -50, -40}
	expected := -10
	min, ok := minmax.Max(list...)

	if min != expected {
		t.Errorf("Expected: %d, Got: %d", expected, min)
	}
	if !ok {
		t.Errorf("Expected: %t, Got: %t", true, ok)
	}
}

func TestMaxPositiveElements(t *testing.T) {
	list := []int{10, 20, 60, 50, 40}
	expected := 60
	min, ok := minmax.Max(list...)

	if min != expected {
		t.Errorf("Expected: %d, Got: %d", expected, min)
	}
	if !ok {
		t.Errorf("Expected: %t, Got: %t", true, ok)
	}
}

func TestMaxMixedElements(t *testing.T) {
	list := []int{80, -40, -50, 10, 30, 100}
	expected := 100
	min, ok := minmax.Max(list...)

	if min != expected {
		t.Errorf("Expected: %d, Got: %d", expected, min)
	}
	if !ok {
		t.Errorf("Expected: %t, Got: %t", true, ok)
	}
}
