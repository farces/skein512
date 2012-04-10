package skein
/* A fairly direct port of the Scala Skein512-512 implementation at
   http://www.h2database.com/skein/

   Hash returns an int8 (as does the original), needs to be cast to
   a byte/uint8 to be usable in a human-readable form using the hex
   package.

   Usage:
       x := skein.NewSkein512()
       r := x.Hash([]byte{bytes,of,message})
*/

type Skein512 struct {
    c []int64
}

func NewSkein512() *Skein512 {
    return &Skein512{c:make([]int64,9)}
}

var d []byte = []byte("ND3EAJ.;1QDGLXV)G>B8-1*R9=GK(6XC")

func (t *Skein512) Hash(data []byte) []int8 {
    t.c = make([]int64,9) //reset c (final blocked vals)

    co1 := uint(56)

    h := make([]int8,64) //return value

    var b0 []byte = []byte("SHA3\x01\x00\x00\x00\x00\x02")
    b := make([]byte,64)
    copy(b,b0)

    t.block(b, 32, int64(196 << co1),0)
    t0 := 0
    t1 := int64(112 << co1)
    for t0 < (len(data) - 64) {
        t0 += 64
        t.block(data, t0, t1, t0-64)
        t1 = int64(48 << co1)
    }
    for i := 0; i<64; i++ {
        if (i + t0) < len(data) {
            b[i] = data[i+t0]
        } else {
            b[i] = 0
        }
    }
    t.block(b, len(data),t1 | (128 << co1), 0)
    t.block(make([]byte,64),8,int64(255 << co1), 0)
    for i := 0; i<64;i++ {
        h[i] = int8(t.c[i/8] >> uint((i & 7) * 8))
    }
    return h
}

func (t *Skein512) block(b []byte, t0 int, t1 int64, o int) {
    x := make([]int64,8)
    k := make([]int64,8)
    t.c[8] = 0x1BD11BDAA9FC1A22 //constant C240
    for i := 0; i<8; i++ {
        for j := 0; j<8; j++ {
            k[i] = (k[i] << 8) + int64(b[o + i * 8 + 7 - j] & 255)
        }
        x[i] = k[i] + t.c[i]
        t.c[8] ^= t.c[i]
    }
    x[5] += int64(t0)
    x[6] += int64(t1)
    y := []int64{int64(t0),int64(t1),int64(t0)^t1}
    for r := 0; r<18; r++ {
        for i := 0;i<16;i++ {
            m := 2 * ((i + (1 + i + i) * (i / 4)) & 3)
            n := (i + 1 + i) & 7
            s := d[16 * (r & 1) + i] - 32
            x[m] += x[n]
            //uint64(int64 value) >> x is equiv. to unsigned right shift
            var temp1 uint64 = uint64(x[n])
            x[n] = ((x[n] << s) | int64(temp1 >> (64 - s))) ^ x[m]
        }
        for i := 0;i<8;i++ {
            x[i] += t.c[(r + 1 + i) % 9]
        }
        x[5] += y[(r + 1) % 3]
        x[6] += y[(r + 2) % 3]
        x[7] += int64(r + 1)
    }
    for i := 0;i<8;i++ {
        t.c[i] = k[i] ^ x[i]
    }
}
