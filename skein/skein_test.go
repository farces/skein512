package skein_test

import (
	"github.com/farces/skein512/skein"
	"testing"
)

func TestSkein512Hash1(t *testing.T) {
	/* Test that 0xFF hashed is equal to the expected int8 array */
	var expected []int8 = []int8{113, -73, -68, -26, -2, 100, 82, 34, 123, -100, -19, 96,
		20, 36, -98, 91, -7, -87, 117, 76, 58, -42, 24, -52, -60, -32, -86, -31, 107, 49,
		108, -56, -54, 105, -115, -122, 67, 7, -19, 62, -128, -74, -17, 21, 112, -127, 42,
		-59, 39, 45, -60, 9, -75, -96, 18, -33, 42, 87, -111, 2, -13, 64, 97, 122}

	x := skein.NewSkein512()
	r := x.Hash([]byte{0xFF})

	for i := 0; i < 64; i++ {
		if r[i] != expected[i] {
			t.Errorf("Element r[%v] was %v, expected %v", i, r[i], expected[i])
		}
	}
}

func TestSkein512Hash2(t *testing.T) {
	/* Test that a string > 512 bits is still hashed correctly */
	input := "aaaaaaaaaabbbbbbbbbbccccccccccddddddddddeeeeeeeeeeffffffffffgggggggggghhhhhhhhhh"
	var expected []int8 = []int8{-35, -28, 52, 111, 42, 114, 52, 63, -45, 78, 14, -78, -113,
		6, -53, 110, 5, -29, 9, -12, -19, -107, 97, -104, 56, 59, 90, 42, -82, 19, 96, -110,
		3, 31, -66, -37, -10, -14, 100, 120, 56, 84, -93, -64, 87, 57, 3, -37, -40, 23, 72,
		61, -15, 91, -61, 108, 19, -25, 95, 59, 54, -98, -84, -61}

	x := skein.NewSkein512()
	r := x.Hash([]byte(input))
	for i := 0; i < 64; i++ {
		if r[i] != expected[i] {
			t.Errorf("Element r[%v] was %v, expected %v", i, r[i], expected[i])
		}
	}
}

func TestSkein512MultipleHash(t *testing.T) {
	/* Test that reusing a Skein512 correctly gives hashes */
	in1 := "ab"
	var expected1 []int8 = []int8{22, -81, -34, 102, 93, -74, -30, 15, 46, 99, 112, 61, -106,
		-67, 42, -63, 82, 9, 76, 59, 59, 36, 35, 98, 53, 55, 21, -20, -58, -21, 110, -71, -16,
		-98, 37, 112, 95, -39, 103, 5, 19, -83, -58, 82, 61, 67, 33, -34, -115, 79, 27, 99, 0,
		84, -100, 125, 117, -65, -67, -87, -110, -52, -78, 112}

	in2 := "cd"
	var expected2 []int8 = []int8{82, -51, 116, -115, 65, 52, 16, 44, 75, -26, -56, 5, 11, -90,
		-33, 44, -96, -98, 42, -3, -16, 86, 120, 44, -105, -84, -41, -120, -41, 77, -36, 114,
		-33, 27, -89, 105, -17, 113, -60, -77, 87, -111, 82, -81, 84, 30, -80, 61, 64, -123, 78,
		-62, 28, -104, 76, -24, 76, 75, 67, 100, 43, 39, 97, -72}

	x := skein.NewSkein512()
	r := x.Hash([]byte(in1))
	s := x.Hash([]byte(in2))

	for i := 0; i < 64; i++ {
		if r[i] != expected1[i] {
			t.Errorf("In HASH 1: Element r[%v] was %v, expected %v", i, r[i], expected1[i])
		}
		if s[i] != expected2[i] {
			t.Errorf("In HASH 2: Element s[%v] was %v, expected %v", i, s[i], expected2[i])
		}
	}
}
