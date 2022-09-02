package e

import (
	"fmt"
)

func Wrapper(msg string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", msg, err)
}
