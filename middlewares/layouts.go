package middlewares

import (
	"github.com/kobeld/duoerl/handlers"
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerl/services"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/mangotemplate"
	"html/template"
)

const (
	TEMPLATE_GLOB_PATTERN = "templates/*/*.html"
	MAIN_LAYOUT           = "main"
)

var (
	preloadedTemplate *template.Template
)

type Header struct {
	AssetsVersion int
	CurrentUser   *users.User
}

type MangoTemplateProvider struct{}

func (h *MangoTemplateProvider) LayoutData(env Env) interface{} {

	header := &Header{
		CurrentUser: services.FetchUserFromEnv(env),
	}

	return header
}

func GetTemplate() *template.Template {
	if preloadedTemplate == nil {
		preloadedTemplate = template.New("")
		mangotemplate.Funcs(preloadedTemplate, handlers.FuncMap)
		template.Must(mangotemplate.ParseGlob(preloadedTemplate, TEMPLATE_GLOB_PATTERN))
	}

	return preloadedTemplate
}

func ProduceLayout(name string) Middleware {
	return mangotemplate.MakeLayout(GetTemplate(), name, &MangoTemplateProvider{})
}

func ProduceRenderer() Middleware {
	return mangotemplate.MakeRenderer(GetTemplate())
}
