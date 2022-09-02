package e

import (
	"fmt"
	"testing"
)

func TestWrapper(t *testing.T) {
	type args struct {
		msg string
		err error
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Error wrapping test, if error is not nil.",
			args: args{
				msg: "Msg",
				err: fmt.Errorf("%s", "error msg"),
			},
			wantErr: true,
		},
		{
			name: "Error wrapping test, if error is nil.",
			args: args{
				msg: "Msg",
				err: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Wrapper(tt.args.msg, tt.args.err); (err != nil) != tt.wantErr {
				t.Errorf("WrapIfErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Benchmark_Wrapper(b *testing.B) {
	errr := fmt.Errorf("%s", "error msg")
	for i := 0; i < b.N; i++ {
		_ = Wrapper("Msg", errr) // calling the tested function
	}
}
