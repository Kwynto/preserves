package preserves

import (
	"bytes"
)

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
