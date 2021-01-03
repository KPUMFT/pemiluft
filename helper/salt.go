package helper

import (
	"io/ioutil"
	"strings"
)

var salt string
var initErr error

// Dalam inisialisasi ini, nilai salt akan diambil dari file yang bernama `.salt`
func init() {
	bSalt, err := ioutil.ReadFile("./.salt")
	if err != nil {
		initErr = err
		return
	}
	salt = strings.TrimSpace(string(bSalt))
}

func salted(s string) string {
	if initErr != nil {
		panic(initErr)
	}
	return s + salt
}
