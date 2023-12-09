// sh.go
package builtins

import (
	"fmt"
	"io"
)

var (
	ErrShInvalidArgCount = fmt.Errorf("%w: unsupported option or argument", ErrInvalidArgCount)
)

// ShBuiltin is a sample sh shell builtin that can invoke sh functions
func ShBuiltin(w io.Writer, args ...string) error {
	if len(args) == 0 {
		// If no arguments are provided, print an error message
		fmt.Fprintln(w, "Error: Function name not specified.")
		return nil
	}

	// For simplicity, let's assume the only supported function is "my_function"
	switch args[0] {
	case "my_function":
		// Invoke the MyFunction with the remaining arguments
		return MyFunction(w, args[1:]...)
	default:
		// If the specified function is not recognized, print an error message
		fmt.Fprintf(w, "Error: Function '%s' not found.\n", args[0])
		return nil
	}
}

// MyFunction is a sample sh function that prints a greeting
func MyFunction(w io.Writer, args ...string) error {
	if len(args) != 1 {
		return ErrShInvalidArgCount
	}

	fmt.Fprintf(w, "Hello, %s!\n", args[0])
	return nil
}
