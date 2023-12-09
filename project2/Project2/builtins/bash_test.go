// bash_test.go
package builtins_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/Freezer105/4600/Project2/builtins"
)

func TestBashBuiltin(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		wantErr error
	}{
		{
			name: "invoke my_function with one argument",
			args: args{
				args: []string{"my_function", "John"},
			},
			wantOut: "Hello, John!\n",
		},
		{
			name: "invoke my_function with multiple arguments",
			args: args{
				args: []string{"my_function", "John", "Doe"},
			},
			wantErr: builtins.ErrBashInvalidArgCount,
		},
		{
			name: "invoke unknown function",
			args: args{
				args: []string{"unknown_function"},
			},
			wantOut: "Error: Function 'unknown_function' not found.\n",
		},
		{
			name: "no function name specified",
			args: args{
				args: []string{},
			},
			wantOut: "Error: Function name not specified.\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// test
			var out bytes.Buffer
			if err := builtins.BashBuiltin(&out, tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("BashBuiltin() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("BashBuiltin() unexpected error: %v", err)
			}
			if got := out.String(); got != tt.wantOut {
				t.Errorf("BashBuiltin() got = %v, want %v", got, tt.wantOut)
			}
		})
	}
}
