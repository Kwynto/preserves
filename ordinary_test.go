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

// func Test_OpenDB(t *testing.T) {
// 	type args struct {
// 		dsn string
// 	}

// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantErr bool
// 	}{
// 		{
// 			name: "No connected to DB MySQL",
// 			args: args{
// 				dsn: "_",
// 			},
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := OpenDB(tt.args.dsn)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("openDB() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			gotType := reflect.TypeOf((*sql.DB)(nil))
// 			if (err != nil) && (reflect.TypeOf(got) != gotType) {
// 				t.Errorf("openDB() error = %v", err)
// 				return
// 			}
// 		})
// 	}
// }

// ------------
// Benchmarking
// ------------

func Benchmark_RandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandInt(0, int64(i+1)) // calling the tested function
	}
}

// func Benchmark_OpenDB(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		_, _ = OpenDB("_") // calling the tested function
// 	}
// }
