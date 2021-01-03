package helper

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strings"
)

// NimToEmail mengkonversikan data NIM menjadi email dengan menambahkan
// '@student.trunojoyo.ac.id'
func NimToEmail(s string) (string, error) {
	return s + "@student.trunojoyo.ac.id", nil
}

// NimToEmailSaltedHash mengkonversikan data NIM menjadi email dengan
// menambahkan '@student.trunojoyo.ac.id', kemudian ditambahi salt dan di-hash.
func NimToEmailSaltedHash(s string) (string, error) {
	nim, err := NimToEmail(s)
	if err != nil {
		return "", err
	}

	return SaltedHash(nim)
}

// JenisEmail memberikan return value "student" jika dan hanya jika email
// menggunakan domain "student.trunojoyo.ac.id". Selain itu akan memberikan
// return value "bukan student".
func JenisEmail(s string) (string, error) {
	idx := strings.Index(s, "@")
	if idx == -1 {
		return "", errors.New("character '@' not found")
	}

	if s[idx+1:] == "student.trunojoyo.ac.id" {
		return "student", nil
	}

	return "bukan student", nil
}

// SaltedHash akan memberikan hasil berupa nilai v yang ditambahi salt kemudian di-hash.
func SaltedHash(v string) (string, error) {
	bSalted := []byte(salted(v))
	bSum := md5.Sum(bSalted)
	return fmt.Sprintf("%x", bSum), nil
}
