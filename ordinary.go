package preserves

import (
	cryptoRand "crypto/rand"
	"errors"
	"math/big"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
)

// The RandInt() function generates a real random integer.
func RandInt(min, max int64) int64 {
	nBig, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(max-min))
	return nBig.Int64() + min
}

// The FindEmail() function looks for an email address in a string and returns it or an error.
func FindEmail(input string) (string, error) {
	emailRegexp := regexp.MustCompile(`.+@.+\..+`)
	first := emailRegexp.FindString(input)
	if first == "" {
		return "", errors.New("email address not found")
	}
	return first, nil
}

// func OpenDB(dsn string) (*sql.DB, error) {
// 	db, err := sql.Open("mysql", dsn)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = db.Ping(); err != nil {
// 		return nil, err
// 	}
// 	db.SetConnMaxLifetime(time.Minute * 5)
// 	db.SetMaxOpenConns(100)
// 	db.SetMaxIdleConns(5)
// 	return db, nil
// }
