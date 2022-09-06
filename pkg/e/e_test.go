package e

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestWrapper(t *testing.T) {
	type args struct {
		msg string
		err error
		wr  *log.Logger
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
				wr:  nil,
			},
			wantErr: true,
		},
		{
			name: "Error wrapping test, if error is nil.",
			args: args{
				msg: "Msg",
				err: nil,
				wr:  nil,
			},
			wantErr: false,
		},
		{
			name: "Error logging.",
			args: args{
				msg: "This is not an error, this diagnostic message is part of a test",
				err: fmt.Errorf("%s", "test msg"),
				wr:  log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Wrapper(tt.args.msg, tt.args.err, tt.args.wr); (err != nil) != tt.wantErr {
				t.Errorf("WrapIfErr() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Benchmark_Wrapper(b *testing.B) {
	errr := fmt.Errorf("%s", "error msg")
	for i := 0; i < b.N; i++ {
		_ = Wrapper("Msg", errr, nil) // calling the tested function
	}
}
