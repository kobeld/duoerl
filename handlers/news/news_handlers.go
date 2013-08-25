package news

import (
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"github.com/theplant/formdata"
	"net/http"
)

var (
	newsFields = []string{"Id", "BrandId", "Title", "Content"}
)

type NewsViewData struct {
	ApiNews   *duoerlapi.News
	NewsInput *duoerlapi.NewsInput
	ApiBrands []*duoerlapi.Brand
	Validated *govalidations.Validated
}

func New(env Env) (status Status, headers Headers, body Body) {

	newsInput := services.NewNews()
	apiBrands, err := services.AllBrands()
	if err != nil {
		panic(err)
	}

	newsViewData := &NewsViewData{
		NewsInput: newsInput,
		ApiBrands: apiBrands,
	}

	mangotemplate.ForRender(env, "news/new", newsViewData)
	return
}

func Create(env Env) (status Status, headers Headers, body Body) {

	newsInput := new(duoerlapi.NewsInput)
	formdata.UnmarshalByNames(env.Request().Request, &newsInput, newsFields)
	newsInput.AuthorId = services.FetchUserIdFromSession(env)

	result, err := services.CreateNews(newsInput)
	if validated, ok := err.(*govalidations.Validated); ok {
		viewData := &NewsViewData{
			NewsInput: newsInput,
			Validated: validated,
		}
		mangotemplate.ForRender(env, "news/new", viewData)
		return
	}
	if err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/news/"+result.Id)
}

func Show(env Env) (status Status, headers Headers, body Body) {
	newsId := env.Request().URL.Query().Get(":id")
	currentUserId := services.FetchUserIdFromSession(env)

	apiNews, err := services.ShowNews(newsId, currentUserId)
	if err != nil {
		panic(err)
	}

	newsViewData := &NewsViewData{
		ApiNews: apiNews,
	}

	mangotemplate.ForRender(env, "news/show", newsViewData)
	return
}

func Edit(env Env) (status Status, headers Headers, body Body) {
	newsId := env.Request().URL.Query().Get(":id")
	currentUser := services.FetchUserFromEnv(env)

	newsInput, err := services.EditNews(currentUser, newsId)
	if err != nil {
		panic(err)
	}

	apiBrands, err := services.AllBrands()
	if err != nil {
		panic(err)
	}

	newsViewData := &NewsViewData{
		NewsInput: newsInput,
		ApiBrands: apiBrands,
	}

	mangotemplate.ForRender(env, "news/edit", newsViewData)
	return
}

func Update(env Env) (status Status, headers Headers, body Body) {
	newsInput := new(duoerlapi.NewsInput)
	formdata.UnmarshalByNames(env.Request().Request, &newsInput, newsFields)

	result, err := services.UpdateNews(newsInput)
	if validated, ok := err.(*govalidations.Validated); ok {
		viewData := &NewsViewData{
			NewsInput: newsInput,
			Validated: validated,
		}
		mangotemplate.ForRender(env, "news/edit", viewData)
		return
	}

	if err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/news/"+result.Id)
}
