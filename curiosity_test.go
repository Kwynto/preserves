package preserves

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// -------
// Testing
// -------

func Test_GenerateId(t *testing.T) {
	testVar := make(map[int]string)
	for i := 0; i < 1000; i++ {
		testVar[i] = GenerateId() // calling the tested function
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
			t.Error("GenerateId() = Error generating unique identifier.")
		}
	}
}

func Test_MeanValue(t *testing.T) {
	type args struct {
		num []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "MeanValue 10.0 10.0",
			args: args{
				num: []float64{10.0, 10.0},
			},
			want: 10.0,
		},
		{
			name: "MeanValue 1.0 2.0 3.0 4.0 5.0",
			args: args{
				num: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
			want: 3.0,
		},
		{
			name: "MeanValue 1.0 2.0 3.0 4.0 5.0 10.0 10.0",
			args: args{
				num: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 10.0, 10.0},
			},
			want: 5.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MeanValue(tt.args.num); got != tt.want {
				t.Errorf("MeanValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_MedianValue(t *testing.T) {
	type args struct {
		num []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "MedianValue 10.0 10.0",
			args: args{
				num: []float64{10.0, 10.0},
			},
			want: 10.0,
		},
		{
			name: "MedianValue 1.0 2.0 3.0 4.0 5.0",
			args: args{
				num: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
			want: 3.0,
		},
		{
			name: "MedianValue 1.0 2.0 3.0 4.0 5.0 10.0 10.0",
			args: args{
				num: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 10.0, 10.0},
			},
			want: 4.0,
		},
		{
			name: "MedianValue 2.0 1.0 4.0 5.0 3.0",
			args: args{
				num: []float64{2.0, 1.0, 4.0, 5.0, 3.0},
			},
			want: 3.0,
		},
		{
			name: "MedianValue 10.0 4.0 1.0 10.0 3.0 2.0 5.0",
			args: args{
				num: []float64{10.0, 4.0, 1.0, 10.0, 3.0, 2.0, 5.0},
			},
			want: 4.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MedianValue(tt.args.num); got != tt.want {
				t.Errorf("MedianValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Variance(t *testing.T) {
	type args struct {
		num []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Variance 10.0 10.0",
			args: args{
				num: []float64{10.0, 10.0},
			},
			want: 0.0,
		},
		{
			name: "Variance 1.0 2.0 3.0 4.0 5.0",
			args: args{
				num: []float64{1.0, 2.0, 3.0, 4.0, 5.0},
			},
			want: 2.0,
		},
		{
			name: "Variance 2.0 1.0 4.0 5.0 3.0",
			args: args{
				num: []float64{2.0, 1.0, 4.0, 5.0, 3.0},
			},
			want: 2.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Variance(tt.args.num); got != tt.want {
				t.Errorf("Variance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ConcatBuffer(t *testing.T) {
	type args struct {
		vals []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Connecting strings via buffer 1",
			args: args{
				vals: []string{"This ", "is ", "even ", "more ", "performant "},
			},
			want: "This is even more performant ",
		},
		{
			name: "Connecting strings via buffer 2",
			args: args{
				vals: []string{"Hello", ", ", "World", "!"},
			},
			want: "Hello, World!",
		},
		{
			name: "Connecting strings via buffer 3",
			args: args{
				vals: []string{"Golang ", "is ", "great", "."},
			},
			want: "Golang is great.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatBuffer(tt.args.vals...); got != tt.want {
				t.Errorf("ConcatBuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ConcatCopy(t *testing.T) {
	type args struct {
		len  int
		vals []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Connecting strings via copy 1",
			args: args{
				len:  29,
				vals: []string{"This ", "is ", "even ", "more ", "performant "},
			},
			want: "This is even more performant ",
		},
		{
			name: "Connecting strings via copy 2",
			args: args{
				len:  13,
				vals: []string{"Hello", ", ", "World", "!"},
			},
			want: "Hello, World!",
		},
		{
			name: "Connecting strings via copy 3",
			args: args{
				len:  16,
				vals: []string{"Golang ", "is ", "great", "."},
			},
			want: "Golang is great.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConcatCopy(tt.args.len, tt.args.vals...); got != tt.want {
				t.Errorf("ConcatCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_DeleteCookie(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		DeleteCookie(&w, "cookieName") // calling the tested function
		io.WriteString(w, "<html><head><title>Title</title></head><body>Body</body></html>")
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	cookie := &http.Cookie{
		Name:   "cookieName",
		Value:  "1234567890",
		MaxAge: 0,
	}
	r.AddCookie(cookie)
	handler(w, r)

	status := w.Code
	// work check
	if status != http.StatusOK {
		t.Errorf("Handler returned %v", status)
	}

	cookies := w.Result().Cookies()
	noErr := true
	for _, v := range cookies {
		if v.Name == "cookieName" && v.Value == "1234567890" {
			noErr = false
		}
	}
	// work check
	if !noErr {
		t.Error("the server did not delete the cookie")
	}
}

func Test_PerformanceTest(t *testing.T) {
	count := PerformanceTest() // calling the tested function
	// work check
	if count < 1 {
		t.Error("PerformanceTest() = No operation was performed.")
	}
	// fmt.Println(count)
}

func Test_UnZipFile(t *testing.T) {
	type args struct {
		zipFile string
		srcFile string
		dstFile string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Правильный тест",
			args: args{
				zipFile: "./testdata/main.zip",
				srcFile: "awesome-go-main/README.md",
				dstFile: "./testdata/unzipedRM1.md",
			},
			want: true,
		},
		{
			name: "Не правильный архив",
			args: args{
				zipFile: "./testdata/master.zip",
				srcFile: "awesome-go-main/README.md",
				dstFile: "./testdata/uzerr1.md",
			},
			want: false,
		},
		{
			name: "Не правильный искомый файл",
			args: args{
				zipFile: "./testdata/main.zip",
				srcFile: "awesome-go-main/ERROR.md",
				dstFile: "./testdata/ERROR.md",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UnZipFile(tt.args.zipFile, tt.args.srcFile, tt.args.dstFile); got != tt.want {
				t.Errorf("UnZipFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

// ------------
// Benchmarking
// ------------

func Benchmark_GenerateId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateId() // calling the tested function
	}
}

func Benchmark_MeanValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MeanValue([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 10.0, 45.0}) // calling the tested function
	}
}

func Benchmark_MedianValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = MedianValue([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 10.0, 45.0}) // calling the tested function
	}
}

func Benchmark_Variance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Variance([]float64{1.0, 2.0, 3.0, 4.0, 5.0, 10.0, 45.0}) // calling the tested function
	}
}

func Benchmark_ConcatBuffer(b *testing.B) {
	var input = []string{"This ", "is ", "even ", "more ", "performant "}
	for i := 0; i < b.N; i++ {
		_ = ConcatBuffer(input...) // calling the tested function
	}
}

func Benchmark_ConcatCopy(b *testing.B) {
	var input = []string{"This ", "is ", "even ", "more ", "performant "}
	for i := 0; i < b.N; i++ {
		_ = ConcatCopy(29, input...) // calling the tested function
	}
}

func Benchmark_DeleteCookie(b *testing.B) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		DeleteCookie(&w, "cookieName") // calling the tested function
	}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/", nil)
	cookie := &http.Cookie{
		Name:   "cookieName",
		Value:  "1234567890",
		MaxAge: 0,
	}
	r.AddCookie(cookie)

	for i := 0; i < b.N; i++ {
		handler(w, r)
	}
}

func Benchmark_UnZipFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = UnZipFile("./testdata/main.zip", "awesome-go-main/README.md", "./testdata/unzipedRM2.md") // calling the tested function
	}
}
