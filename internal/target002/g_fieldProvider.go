// Code generated by github.com/frm-adiputra/csv2postgres DO NOT EDIT

package target002

import (
	"github.com/frm-adiputra/csv2postgres/utils"
)

// FieldProvider implements pipeline.FieldProvider interface.
type FieldProvider struct {
	fieldMap map[string]int
}

// Setup sets up the provider, must be called on initialization (before other
// calls to ProvideRow).
func (p *FieldProvider) Setup(header []string) error {
	renameMap := map[string]string{
		"waktu_str":        "Timestamp",
		"nim_email_hashed": "Username",
		"jenis_email":      "Username",
		"persetujuan":      "SAYA MENGGUNAKAN HAK PILIH DALAM PEMILU E-VOTE SECARA SADAR DAN TANPA TEKANAN DARI SIAPAPUN:",
		"dapil":            "PROGRAM STUDI :",
		"pilihan_gubernur": "Silahkan pilih calon Gubernur dan Wakil Gubernur FT UTM",
		"pilihan_dpm":      "Silahkan pilih calon Anggota DPM FT UTM",
		"pilihan_himatro":  "Silahkan pilih calon KETUA UMUM dan WAKIL KETUA UMUM HIMATRO FT UTM ",
		"pilihan_hmti":     "Silahkan pilih calon KETUA UMUM dan WAKIL KETUA UMUM HMTI FT UTM",
	}

	fieldMap, err := utils.CreateHeader(header, renameMap)
	if err != nil {
		return err
	}
	p.fieldMap = fieldMap
	return nil
}

// ProvideField provide a row to be accessed using map of field names.
func (p FieldProvider) ProvideField(row []string) (map[string]interface{}, error) {
	return utils.ProvideField(p.fieldMap, row)
}
