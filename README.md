# Skein512-512 in Go
See http://www.h2database.com/skein/ for the original implementation. Based on the Scala version, which is (most likely) optimized for size. Scala implementation based on the C reference implementation v1.3.<br>

```go
go get github.com/farces/skein512/skein
import "github.com/farces/skein512/skein"
```

## Usage

```go
hasher := skein.NewSkein512()
message := "your message here"
result := hasher.Hash([]byte(message))
```

.Hash(x []byte) takes a byte array rather than a string as it can/will be used to hash non-strings (i.e. file data, etc.)


## Disclaimer
Other than the tests in the test suite (short string, > 512bit string, multiple consecuitive hashes) this package has not been widely tested. Do not use with the impression that it is complete or faultless, do not use in Production Code without significant testing. For educational use only.
