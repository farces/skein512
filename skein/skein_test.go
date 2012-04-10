package skein_test

import (
	"github.com/farces/skein512/skein"
	"testing"
)

func TestSkein512Hash(t *testing.T) {
	/* Test that 0xFF hashed is equal to the expected int8 array */
	var expected []int8 = []int8{113, -73, -68, -26, -2, 100, 82, 34, 123, -100, -19, 96,
		20, 36, -98, 91, -7, -87, 117, 76, 58, -42, 24, -52, -60, -32, -86, -31, 107, 49, 108, -56,
		-54, 105, -115, -122, 67, 7, -19, 62, -128, -74, -17, 21, 112, -127, 42, -59, 39, 45, -60,
		9, -75, -96, 18, -33, 42, 87, -111, 2, -13, 64, 97, 122}

	x := skein.NewSkein512()
	r := x.Hash([]byte{0xFF})

	for i := 0; i < 64; i++ {
		if r[i] != expected[i] {
			t.Errorf("Element r[i] was %v, expected %v", r[i], expected[i])
		}
	}
}
