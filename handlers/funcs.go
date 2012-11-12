package handlers

import (
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"html/template"
)

var FuncMap = template.FuncMap{
	"errorOn": errorOn,
}

type ErrorData struct {
	Errors    govalidations.Errors
	FieldName string
}

func errorOn(validated *govalidations.Validated, fieldName string) (r template.HTML) {
	if validated == nil {
		return
	}

	if !validated.HasError() {
		return
	}
	if !validated.Errors.Has(fieldName) {
		return
	}
	r = template.HTML(mangotemplate.RenderToString("errors", &ErrorData{
		Errors:    validated.Errors,
		FieldName: fieldName,
	}))
	return
}
