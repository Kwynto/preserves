package e

import (
	"fmt"
	"log"
)

// The Wrapper function is needed to catch the error message inside your function.
// To be intercepted, your function must have a named output characteristic of type error, such as `func myFunc() err error`
// For simple hooking, you need to start your function with defer `func() { e.Wrapper("Your msg", err, nil) }`
// Now you can simply return an error like this `return nil` or `return err`, the error will be caught and formatted as needed, including being written to the log
func Wrapper(msg string, err error, wrLog *log.Logger) error {
	if err == nil {
		return nil
	}
	newErr := fmt.Errorf("%s: %w", msg, err)
	if wrLog != nil {
		// There might be a logger here
		wrLog.Print(newErr)
	}
	return newErr
}
