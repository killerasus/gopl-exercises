package rotate_test

import (
	"exerc4_4/rotate"
	"testing"
)

func TestRotateLeft1Time(t *testing.T) {
	array := [...]int{0, 1, 2, 3, 4}
	rotate.Rotate(array[:], 1)

	for i, v := range [...]int{1, 2, 3, 4, 0} {
		if array[i] != v {
			t.Errorf("wanted %d, got %d", v, array[i])
		}
	}
}
