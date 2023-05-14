package ordinary

import (
	"testing"
	// _ "github.com/go-sql-driver/mysql"
)

var resultFibo uint

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

func Test_Fibo(t *testing.T) {
	type args struct {
		num uint
	}
	tests := []struct {
		name string
		args args
		want uint
	}{
		{
			name: "Fibonacci number 0.",
			args: args{
				num: 0,
			},
			want: 0,
		},
		{
			name: "Fibonacci number 1.",
			args: args{
				num: 1,
			},
			want: 1,
		},
		{
			name: "Fibonacci number 2.",
			args: args{
				num: 2,
			},
			want: 1,
		},
		{
			name: "Fibonacci number 3.",
			args: args{
				num: 3,
			},
			want: 2,
		},
		{
			name: "Fibonacci number 4.",
			args: args{
				num: 4,
			},
			want: 3,
		},
		{
			name: "Fibonacci number 5.",
			args: args{
				num: 5,
			},
			want: 5,
		},
		{
			name: "Fibonacci number 6.",
			args: args{
				num: 6,
			},
			want: 8,
		},
		{
			name: "Fibonacci number 7.",
			args: args{
				num: 7,
			},
			want: 13,
		},
		{
			name: "Fibonacci number 8.",
			args: args{
				num: 8,
			},
			want: 21,
		},
		{
			name: "Fibonacci number 9.",
			args: args{
				num: 9,
			},
			want: 34,
		},
		{
			name: "Fibonacci number 10.",
			args: args{
				num: 10,
			},
			want: 55,
		},
		{
			name: "Fibonacci number 11.",
			args: args{
				num: 11,
			},
			want: 89,
		},
		{
			name: "Fibonacci number 12.",
			args: args{
				num: 12,
			},
			want: 144,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fibo(tt.args.num); got != tt.want {
				t.Errorf("Fibo() = %v, want %v", got, tt.want)
			}
		})
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

func Test_DownloadFile(t *testing.T) {
	type args struct {
		source string
		dst    string
	}
	tests := []struct {
		name    string
		args    args
		wanterr bool
	}{
		{
			name: "Download test - Source error",
			args: args{
				source: "ksh3khs87&kdhf9*&fkshd97kjk^(_@E.,dsj989sd||}NF",
				dst:    "./testdata/",
			},
			wanterr: true,
		},
		{
			name: "Download test - Normal test",
			args: args{
				source: "https://github.com/Kwynto/Kwynto/README.md",
				dst:    "./testdata/",
			},
			wanterr: false,
		},
		{
			name: "Download test - Space test",
			args: args{
				source: "",
				dst:    "./testdata/",
			},
			wanterr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := DownloadFile(tt.args.source, tt.args.dst); err != nil && tt.wanterr == false {
				t.Errorf("DownloadFile() recive error = %v", err)
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

// Helper function for Fibo() benchmark
func benchmarkFibo(b *testing.B, n uint) {
	var r uint
	for i := 0; i < b.N; i++ {
		r = Fibo(n) // calling the tested function
	}
	resultFibo = r
}

func Benchmark_30Fibo(b *testing.B) {
	benchmarkFibo(b, 30) // helper function call
}

func Benchmark_50Fibo(b *testing.B) {
	benchmarkFibo(b, 50) // helper function call
}

func Benchmark_150Fibo(b *testing.B) {
	benchmarkFibo(b, 150) // helper function call
}

func Benchmark_FindEmail(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = FindEmail("email@example.com") // calling the tested function
	}
}

func Benchmark_DownloadFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = DownloadFile("", "./testdata/") // calling the tested function
	}
}
