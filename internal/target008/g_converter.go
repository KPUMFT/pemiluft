// Code generated by github.com/frm-adiputra/csv2postgres DO NOT EDIT

package target008

import (
	"github.com/frm-adiputra/csv2postgres/utils"
)

// Converter implements pipeline.Converter interface.
type Converter struct{}

// Convert converts field's values.
func (c Converter) Convert(fields map[string]interface{}) (map[string]interface{}, error) {
	if fields["dapil"] == "" {
		return nil, utils.ErrEmptyValue("dapil")
	}
	vDapil := fields["dapil"]
	fields["dapil"] = vDapil

	return fields, nil
}
