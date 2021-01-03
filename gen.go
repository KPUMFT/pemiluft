// +build ignore

package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/frm-adiputra/csv2postgres"
)

// fungsi ini akan dieksekusi melalui perintah `go generate`.
func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "salt" {
			if err := generateSalt(); err != nil {
				exitWithError(err)
			}
		}
	} else {
		if err := generateDDLCode(); err != nil {
			exitWithError(err)
		}
	}
}

// generateDDLCode menggenerate kode program yang akan digunakan untuk
// meng-import data ke database.
func generateDDLCode() error {
	g := csv2postgres.Generator{
		BaseImportPath: "evote",
		DefaultSchema:  "evote",
	}
	return g.Generate()
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

// generateSalt menggenerate salt secara acak dan menyimpannya ke dalam file
// `.salt`.
func generateSalt() error {
	n, err := randomNum()
	if err != nil {
		return err
	}

	d1 := []byte(n)
	err = ioutil.WriteFile("./.salt", d1, 0600)
	if err != nil {
		return err
	}

	return nil
}

// randomNum menggenerate nilai acak yang akan digunakan sebagai salt.
func randomNum() (string, error) {
	n, err := rand.Prime(rand.Reader, 64)
	if err != nil {
		return "", err
	}
	return n.String(), nil
}
