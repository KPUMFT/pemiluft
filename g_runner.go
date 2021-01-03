// Code generated by github.com/frm-adiputra/csv2postgres DO NOT EDIT

package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/frm-adiputra/csv2postgres/pipeline"
	_ "github.com/lib/pq"
)

var cfg pipeline.DBConfig

func init() {
	var err error
	cfg, err = pipeline.NewDBConfig("./db.yaml")
	if err != nil {
		exitWithError(err)
	}
}

func newDBConn() (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnectionString())
	if err != nil {
		return nil, err
	}
	return db, nil
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}