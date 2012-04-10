package main

import (
	"encoding/hex"
	"fmt"
	"github.com/farces/skein512/skein"
)

func main() {
	x := skein.NewSkein512()
	r := x.Hash([]byte{0xFF})
	out := make([]byte, 64)

	for i := 0; i < 64; i++ {
		out[i] = byte(r[i])
	}

	fmt.Println(hex.EncodeToString(out))
}
