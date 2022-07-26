package preserves

import (
	"archive/zip"
	"bytes"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"time"
)

// Math

// The GenerateId() generates a new id in a random, cryptographically secure manner
func GenerateId() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// The MeanValue() function returns the average value from a slice of real numbers.
func MeanValue(x []float64) float64 {
	sum := float64(0)
	for _, v := range x {
		sum = sum + v
	}
	return sum / float64(len(x))
}

// The MedianValue() function returns the median value of a slice of real numbers.
func MedianValue(x []float64) float64 {
	sort.Float64s(x)
	length := len(x)
	if length%2 == 1 {
		return x[(length-1)/2]
	} else {
		return (x[length/2] + x[(length/2)-1]) / 2
	}
}

// The Variance() function returns the variance from a slice of real numbers.
func Variance(x []float64) float64 {
	mean := MeanValue(x)
	sum := float64(0)
	for _, v := range x {
		sum = sum + (v-mean)*(v-mean)
	}
	return sum / float64(len(x))
}

// Strings

// The ConcatBuffer() function quickly concatenates strings.
// This feature is useful for optimizing code that needs to quickly concatenate strings.
func ConcatBuffer(vals ...string) string {
	buffer := bytes.Buffer{}
	for _, val := range vals {
		buffer.WriteString(val)
	}
	return buffer.String()
}

// The ConcatCopy() function is a very fast string concatenation, but requires knowing the length of the string, since anything over that length will be lost.
// This feature is useful for optimizing code that needs to quickly concatenate strings.
func ConcatCopy(len int, vals ...string) string {
	bs := make([]byte, len)
	bl := 0
	for _, val := range vals {
		bl += copy(bs[bl:], []byte(val))
	}
	return string(bs[:])
}

// Web

// The DeleteCookie(w) function deletes the cookie
func DeleteCookie(w *http.ResponseWriter, cookieName string) {
	cookie := &http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(*w, cookie)
}

// Other

// The PerformanceTest() function measures the amount of time the current time is added to the slice in 3 seconds.
// This feature can be used to evaluate the performance of your computer.
func PerformanceTest() int {
	to := time.After(3 * time.Second)
	list := make([]string, 0)
	done := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-to:
				done <- true
				return
			default:
				list = append(list, time.Now().String())
			}
		}
	}()

	<-done
	return len(list)
}

// Функция UnZipFile() распаковывает конкретный файл из архива и записывает в указанное место. // FIXME: перевести
func UnZipFile(zipFile, srcFile, dstFile string) bool {
	var status bool = false
	zipR, err := zip.OpenReader(zipFile)
	if err != nil {
		return status
	}
	for _, file := range zipR.File {
		if file.Name == srcFile {
			r, _ := file.Open()
			outF, _ := os.Create(dstFile)
			_, _ = io.Copy(outF, r)
			_ = r.Close()
			status = true
			break
		}
	}
	return status
}
