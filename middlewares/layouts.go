package middlewares

import (
	"github.com/kobeld/duoerl/configs"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/mangotemplate"
	"html/template"
)

const (
	TEMPLATES   = "templates/*/*.html"
	MAIN_LAYOUT = "main"
)

type Header struct {
	AssetsVersion int
}

type provider struct{}

func (h *provider) LayoutData(env Env) interface{} {

	header := &Header{
		AssetsVersion: configs.AssetsVersion,
	}

	return header
}

func ProduceLayout(name string) Middleware {
	tpl := template.New("")
	template.Must(tpl.ParseGlob(TEMPLATES))
	return mangotemplate.MakeLayout(tpl, name, &provider{})
}

func ProduceRenderer() Middleware {
	tpl := template.New("")
	template.Must(tpl.ParseGlob(TEMPLATES))
	return mangotemplate.MakeRenderer(tpl)
}
