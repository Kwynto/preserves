package preserves

import (
	"testing"
	// _ "github.com/go-sql-driver/mysql"
)

// -------
// Testing
// -------

func Test_RandInt(t *testing.T) {
	testVar := make(map[int]int64)
	for i := 0; i < 10; i++ {
		testVar[i] = RandInt(0, 10000000) // calling the tested function
	}
	for _, v1 := range testVar {
		count := 0
		for _, v2 := range testVar {
			if v1 == v2 {
				count++
			}
		}
		// work check
		if count > 1 {
			t.Error("RandInt() = Error generating unique random integer.")
		}
	}
}

func Test_FindEmail(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Email search 1",
			args: args{
				val: "email@example.com",
			},
			want: "email@example.com",
		},
		{
			name: "Email search 2",
			args: args{
				val: "other@domain.com",
			},
			want: "other@domain.com",
		},
		{
			name: "Email search 3",
			args: args{
				val: "other-mail@other-domain.xyz",
			},
			want: "other-mail@other-domain.xyz",
		},
		{
			name: "Email search 4",
			args: args{
				val: "other-mail&@other-!(domain)",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := FindEmail(tt.args.val); got != tt.want {
				t.Errorf("FindEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ------------
// Benchmarking
// ------------

func Benchmark_RandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandInt(0, int64(i+1)) // calling the tested function
	}
}

func Benchmark_FindEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = FindEmail("email@example.com") // calling the tested function
	}
}
