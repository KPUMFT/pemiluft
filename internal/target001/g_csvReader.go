// Code generated by github.com/frm-adiputra/csv2postgres DO NOT EDIT

package target001

import "github.com/frm-adiputra/csv2postgres/pipeline"

// NewCSVReader creates a new CSVReader.
func NewCSVReader() *pipeline.CSVReader {
	return &pipeline.CSVReader{
		FileName:  "data/daftar-pemilih.csv",
		Separator: ',',
	}
}
