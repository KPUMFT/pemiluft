// Code generated by github.com/frm-adiputra/csv2postgres DO NOT EDIT

package target005

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
		"calon": "calon",
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