package preserves

import (
	cryptoRand "crypto/rand"
	"errors"
	"math/big"
	"regexp"
	"time"
)

// Math

// The RandInt() function generates a real random integer.
func RandInt(min, max int64) int64 {
	nBig, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(max-min))
	return nBig.Int64() + min
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

// Other

// The PerformanceTest() function measures the amount of time the current time is added to the slice in 3 seconds.
// This feature can be used to evaluate the performance of your computer.
func PerformanceTest() int {
	to := time.After(3 * time.Second)
	list := make([]string, 0)
	done := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-to:
				done <- true
				return
			default:
				list = append(list, time.Now().String())
			}
		}
	}()

	<-done
	return len(list)
}
