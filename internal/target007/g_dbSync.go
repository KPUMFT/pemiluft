// Code generated by github.com/frm-adiputra/csv2postgres DO NOT EDIT

package target007

import (
	"database/sql"
	"fmt"
	"io"

	"github.com/frm-adiputra/csv2postgres/pipeline"
	"github.com/lib/pq"
)

// DBSynchronizer implements pipeline.DBSynchronizer.
type DBSynchronizer struct {
	name         string
	sqlCreate    string
	sqlDelete    string
	sqlDrop      string
	dependsOn    []string
	recordReader *pipeline.RecordReader
}

// NewDBSynchronizer creates a new instance.
func NewDBSynchronizer() DBSynchronizer {
	r := &pipeline.RecordReader{
		Name:          "ref_dapil",
		RowReader:     NewCSVReader(),
		FieldProvider: &FieldProvider{},
		Converter:     Converter{},
		Computer:      Computer{},
		Validator:     Validator{},
	}

	return DBSynchronizer{
		name:         "evote.ref_dapil",
		dependsOn:    []string{},
		recordReader: r,
		sqlDelete:    `DELETE FROM "evote"."ref_dapil"`,
		sqlDrop:      `DROP TABLE IF EXISTS "evote"."ref_dapil"`,
		sqlCreate: `
            CREATE TABLE "evote"."ref_dapil" (
                "dapil" text NOT NULL,
                PRIMARY KEY (dapil)
            )
        `,
	}
}

// Name returns the table's name
func (d DBSynchronizer) Name() string {
	return d.name
}

// RowCount returns number of rows that is filled.
func (d DBSynchronizer) RowCount() int64 {
	return d.recordReader.RowReader.RowCount()
}

// DependsOn returns other tables that this table depends on
func (d DBSynchronizer) DependsOn() []string {
	return d.dependsOn
}

// Create table
func (d DBSynchronizer) Create(db *sql.DB) error {
	_, err := db.Exec(d.sqlCreate)
	return err
}

// Delete all rows from table
func (d DBSynchronizer) Delete(db *sql.DB) error {
	_, err := db.Exec(d.sqlDelete)
	return err
}

// Drop table
func (d DBSynchronizer) Drop(db *sql.DB) error {
	_, err := db.Exec(d.sqlDrop)
	return err
}

// Fill rows
func (d DBSynchronizer) Fill(db *sql.DB) error {
	err := d.recordReader.Open()
	if err != nil {
		return fmt.Errorf("%s: %w", d.recordReader.Name, err)
	}

	defer d.recordReader.Close()

	txn, err := db.Begin()
	if err != nil {
		return fmt.Errorf("%s: %w", d.recordReader.Name, err)
	}

	stmt, err := txn.Prepare(pq.CopyInSchema(
		"evote",
		"ref_dapil",
		"dapil",
	))
	if err != nil {
		if rollbackErr := txn.Rollback(); rollbackErr != nil {
			return fmt.Errorf("%s: %w", d.recordReader.Name,
				fmt.Errorf("failed to rollback: %s: %w", rollbackErr.Error(), err))
		}
		return fmt.Errorf("%s: %w", d.recordReader.Name, err)
	}

	for {
		rec, err := d.recordReader.ReadRecord()
		if err == io.EOF {
			break
		}
		if err != nil {
			stmt.Exec()
			if rollbackErr := txn.Rollback(); rollbackErr != nil {
				return fmt.Errorf("%s record #%d: %w", d.recordReader.Name,
					d.recordReader.RowReader.RowCount(),
					fmt.Errorf("failed to rollback: %s: %w", rollbackErr.Error(), err))
			}
			return fmt.Errorf("%s record #%d: %w",
				d.recordReader.Name, d.recordReader.RowReader.RowCount(), err)
		}
		_, err = stmt.Exec(
			rec["dapil"],
		)
		if err != nil {
			stmt.Exec()
			if rollbackErr := txn.Rollback(); rollbackErr != nil {
				return fmt.Errorf("%s record #%d: %w", d.recordReader.Name,
					d.recordReader.RowReader.RowCount(),
					fmt.Errorf("failed to rollback: %s: %w", rollbackErr.Error(), err))
			}
			return fmt.Errorf("%s record #%d: %w",
				d.recordReader.Name, d.recordReader.RowReader.RowCount(), err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		if rollbackErr := txn.Rollback(); rollbackErr != nil {
			return fmt.Errorf("%s: %w", d.recordReader.Name,
				fmt.Errorf("failed to rollback: %s: %w", rollbackErr.Error(), err))
		}
		return fmt.Errorf("%s: %w", d.recordReader.Name, err)
	}

	err = stmt.Close()
	if err != nil {
		if rollbackErr := txn.Rollback(); rollbackErr != nil {
			return fmt.Errorf("%s: %w", d.recordReader.Name,
				fmt.Errorf("failed to rollback: %s: %w", rollbackErr.Error(), err))
		}
		return fmt.Errorf("%s: %w", d.recordReader.Name, err)
	}

	err = txn.Commit()
	if err != nil {
		if rollbackErr := txn.Rollback(); rollbackErr != nil {
			return fmt.Errorf("%s: %w", d.recordReader.Name,
				fmt.Errorf("failed to rollback: %s: %w", rollbackErr.Error(), err))
		}
		return fmt.Errorf("%s: %w", d.recordReader.Name, err)
	}
	return nil
}
