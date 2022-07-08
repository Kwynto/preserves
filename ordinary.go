package preserves

import (
	cryptoRand "crypto/rand"
	"math/big"
)

// The RandInt() function generates a real random integer.
func RandInt(min, max int64) int64 {
	nBig, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(max-min))
	return nBig.Int64() + min
}
