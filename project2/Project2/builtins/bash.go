// bash.go
package builtins

import (
	"fmt"
	"io"
)

var (
	ErrBashInvalidArgCount = fmt.Errorf("%w: unsupported option or argument", ErrInvalidArgCount)
)

// BashBuiltin is a sample bash shell builtin that can invoke bash functions
func BashBuiltin(w io.Writer, args ...string) error {
	if len(args) == 0 {
		// If no arguments are provided, print an error message
		fmt.Fprintln(w, "Error: Function name not specified.")
		return nil
	}

	// For simplicity, let's assume the only supported function is "my_function"
	switch args[0] {
	case "my_function":
		// Invoke the MyFunction with the remaining arguments
		return invokeMyFunction(w, args[1:]...)
	default:
		// If the specified function is not recognized, print an error message
		fmt.Fprintf(w, "Error: Function '%s' not found.\n", args[0])
		return nil
	}
}

// invokeMyFunction is a sample bash function that prints a greeting
func invokeMyFunction(w io.Writer, args ...string) error {
	// You can use $1, $2, etc. to access function arguments
	if len(args) != 0 {
		fmt.Fprintf(w, "Hello, %s!\n", args[0])
	} else {
		fmt.Fprintln(w, "Error: Argument not provided.")
	}
	return nil
}
