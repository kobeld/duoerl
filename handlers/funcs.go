package handlers

import (
	"github.com/shaoshing/train"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"html/template"
	"reflect"
)

var FuncMap = template.FuncMap{
	"ErrorOn": errorOn,
	"EqualId": equalId,
	"Equal":   eq,

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

func equalId(idOne, idTwo string) bool {
	if idOne == idTwo {
		return true
	}
	return false
}

// eq reports whether the first argument is equal to
// any of the remaining arguments.
func eq(args ...interface{}) bool {
	if len(args) == 0 {
		return false
	}
	x := args[0]
	switch x := x.(type) {
	case string, int, int64, byte, float32, float64:
		for _, y := range args[1:] {
			if x == y {
				return true
			}
		}
		return false
	}

	for _, y := range args[1:] {
		if reflect.DeepEqual(x, y) {
			return true
		}
	}
	return false
}
