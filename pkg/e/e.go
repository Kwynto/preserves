package e

import (
	"fmt"
)

func Wrapper(msg string, err error) error {
	if err == nil {
		return nil
	}
	newErr := fmt.Errorf("%s: %w", msg, err)
	// There might be a logger here
	return newErr
}
