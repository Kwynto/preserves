package preserves

import (
	"bytes"
	"crypto/rand"
	"fmt"
)

// The GenerateId() generates a new id in a random, cryptographically secure manner
func GenerateId() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// The ConcatBuffer() function quickly concatenates strings.
// This feature is useful for optimizing code that needs to quickly concatenate strings.
func ConcatBuffer(vals ...string) string {
	// strings := []string{"This ", "is ", "even ", "more ", "performant "}
	buffer := bytes.Buffer{}
	for _, val := range vals {
		buffer.WriteString(val)
	}
	return buffer.String()
}
