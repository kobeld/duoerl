package handlers

import (
	"github.com/shaoshing/train"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"html/template"
)

var FuncMap = template.FuncMap{
	"ErrorOn": errorOn,

	"javascript_tag":            train.JavascriptTag,
	"stylesheet_tag":            train.StylesheetTag,
	"stylesheet_tag_with_param": train.StylesheetTagWithParam,
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
