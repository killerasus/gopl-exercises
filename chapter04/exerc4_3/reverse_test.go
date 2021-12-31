package reverse

import (
	"testing"
)

func TestReverseNoItem(t *testing.T) {
	var array [5]int
	want := len(array)

	reverse(&array)

	got := len(array)

	if want != got {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestReverseFiveItems(t *testing.T) {
	var array = [5]int{0, 1, 2, 3, 4}
	reverse(&array)

	if array[0] != 4 {
		t.Errorf("wanted %d, got %d", 4, array[0])
	}

	if array[1] != 3 {
		t.Errorf("wanted %d, got %d", 3, array[1])
	}

	if array[2] != 2 {
		t.Errorf("wanted %d, got %d", 2, array[2])
	}

	if array[3] != 1 {
		t.Errorf("wanted %d, got %d", 1, array[3])
	}

	if array[4] != 0 {
		t.Errorf("wanted %d, got %d", 0, array[4])
	}
}
