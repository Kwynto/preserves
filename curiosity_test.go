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

// ------------
// Benchmarking
// ------------

func Benchmark_GenerateId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenerateId() // calling the tested function
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
