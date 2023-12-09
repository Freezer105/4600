// csh_test.go
package builtins_test

import (
	"bytes"
	"errors"

	//"fmt"
	"os"
	"testing"

	"github.com/Freezer105/4600/Project2/builtins"
)

func TestCshBuiltin(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		setEnv  map[string]string
		wantOut string
		wantErr error
	}{
		{
			name: "invoke my_cd with one argument",
			args: args{
				args: []string{"my_cd", "/path/to/directory"},
			},
			wantErr: nil, // Assuming the directory exists
		},
		{
			name: "invoke my_cd with multiple arguments",
			args: args{
				args: []string{"my_cd", "/path/to/directory", "extra_argument"},
			},
			wantErr: builtins.ErrMyCDInvalidArgCount,
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
			// setup
			for k, v := range tt.setEnv {
				os.Setenv(k, v)
			}

			// test
			var out bytes.Buffer
			if err := builtins.CshBuiltin(&out, tt.args.args...); tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("CshBuiltin() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			} else if err != nil {
				t.Fatalf("CshBuiltin() unexpected error: %v", err)
			}
			if got := out.String(); got != tt.wantOut {
				t.Errorf("CshBuiltin() got = %v, want %v", got, tt.wantOut)
			}
		})
	}
}
