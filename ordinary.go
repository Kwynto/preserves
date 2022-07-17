package preserves

import (
	cryptoRand "crypto/rand"
	mathRand "math/rand"
	"errors"
	"math/big"
	"regexp"
)

// Math

// The RandInt() function generates a real random integer.
func RandInt(min, max int64) int64 {
	// Attention: import cryptoRand "crypto/rand"
	nBig, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(max-min))
	return nBig.Int64() + min
}

// The UsualRandInt() function generates a pseudo-random integer.
func UsualRandInt(min, max int) int {
	// Attention: import mathRand "math/rand"
	return mathRand.Intn(max-min) + min
}

func fiboInternal(n uint, a, b uint) uint {
	// Internal function for use in Fibo(n)
	// This function implements the final recursion.
	if n == 1 {
		return b
	}
	return fiboInternal(n-1, b, a+b)
}

// The Fibo() function is a fast implementation of the Fibonacci number via finite recursion.
func Fibo(n uint) uint {
	if n == 0 {
		return 0
	}
	return fiboInternal(n, 0, 1)
}

// Web

// The FindEmail() function looks for an email address in a string and returns it or an error.
func FindEmail(input string) (string, error) {
	emailRegexp := regexp.MustCompile(`.+@.+\..+`)
	first := emailRegexp.FindString(input)
	if first == "" {
		return "", errors.New("email address not found")
	}
	return first, nil
}
