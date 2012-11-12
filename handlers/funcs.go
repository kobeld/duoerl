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
	Errors    *govalidations.Errors
	FieldName string
}

func errorOn(errors *govalidations.Errors, fieldName string) (r template.HTML) {
	if errors == nil {
		return
	}
	if !errors.Has(fieldName) {
		return
	}
	r = template.HTML(mangotemplate.RenderToString("errors", &ErrorData{
		Errors:    errors,
		FieldName: fieldName,
	}))
	return
}
