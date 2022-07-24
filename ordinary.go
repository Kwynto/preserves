package preserves

import (
	cryptoRand "crypto/rand"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
)

// Math

// The RandInt() function generates a real random integer.
func RandInt(min, max int64) int64 {
	// Attention: import cryptoRand "crypto/rand"
	nBig, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(max-min))
	return nBig.Int64() + min
}

func fiboInternal(n uint, a, b uint) uint {
	// Internal function for use in Fibo(n)
	// This function implements the final recursion.
	if n == 1 {
		return b
	}
	return fiboInternal(n-1, b, a+b)
}

// The Fibo() function is a fast implementation of the Fibonacci number via finite recursion.
func Fibo(n uint) uint {
	if n == 0 {
		return 0
	}
	return fiboInternal(n, 0, 1)
}

// Web

// The FindEmail() function looks for an email address in a string and returns it or an error.
func FindEmail(input string) (string, error) {
	emailRegexp := regexp.MustCompile(`.+@.+\..+`)
	first := emailRegexp.FindString(input)
	if first == "" {
		return "", errors.New("email address not found")
	}
	return first, nil
}

func DownloadFile(sourceUrl string, dstFolder string) (int64, error) {
	// Build fileName from fullPath
	fileURL, _ := url.Parse(sourceUrl)
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := segments[len(segments)-1]

	// Create blank file
	// fileName = ConcatBuffer(dstFolder, fileName)
	fileName = fmt.Sprint(dstFolder, fileName)
	file, err := os.Create(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// Put content on file
	// client := http.Client{
	// 	CheckRedirect: func(r *http.Request, via []*http.Request) error {
	// 		r.URL.Opaque = r.URL.Path
	// 		return nil
	// 	},
	// }
	client := http.Client{}
	resp, err := client.Get(sourceUrl)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)
	return size, err
}
