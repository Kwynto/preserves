package preserves

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"net/http"
	"time"
)

// Math

// The GenerateId() generates a new id in a random, cryptographically secure manner
func GenerateId() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// Strings

// The ConcatBuffer() function quickly concatenates strings.
// This feature is useful for optimizing code that needs to quickly concatenate strings.
func ConcatBuffer(vals ...string) string {
	buffer := bytes.Buffer{}
	for _, val := range vals {
		buffer.WriteString(val)
	}
	return buffer.String()
}

// The ConcatCopy() function is a very fast string concatenation, but requires knowing the length of the string, since anything over that length will be lost.
// This feature is useful for optimizing code that needs to quickly concatenate strings.
func ConcatCopy(len int, vals ...string) string {
	bs := make([]byte, len)
	bl := 0
	for _, val := range vals {
		bl += copy(bs[bl:], []byte(val))
	}
	return string(bs[:])
}

// Web

// The DeleteCookie(w) function deletes the cookie
func DeleteCookie(w *http.ResponseWriter, cookieName string) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(*w, cookie)
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
