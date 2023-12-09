// csh.go
package builtins

import (
	"fmt"
	"io"
	"os"
)

var (
	ErrCshInvalidArgCount  = fmt.Errorf("%w: unsupported option or argument", ErrInvalidArgCount)
	ErrMyCDInvalidArgCount = fmt.Errorf("%w: expected one argument (directory)", ErrInvalidArgCount)
)

// MyCDFunction is a sample function in the C shell that changes the directory
func MyCDFunction(args ...string) error {
	switch len(args) {
	case 1:
		return os.Chdir(args[0])
	default:
		return ErrMyCDInvalidArgCount
	}
}

// CshBuiltin is a sample C shell builtin that can invoke the MyCDFunction or other functions
func CshBuiltin(w io.Writer, args ...string) error {
	if len(args) == 0 {
		// If no arguments are provided, print an error message
		fmt.Fprintln(w, "Error: Function name not specified.")
		return nil
	}

	// For simplicity, let's assume the only supported function is "my_cd"
	switch args[0] {
	case "my_cd":
		// Invoke the MyCDFunction with the remaining arguments
		return MyCDFunction(args[1:]...)
	default:
		// If the specified function is not recognized, print an error message
		fmt.Fprintf(w, "Error: Function '%s' not found.\n", args[0])
		return nil
	}
}
