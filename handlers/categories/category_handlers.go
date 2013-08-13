package categories

import (
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/formdata"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"net/http"
)

var (
	categoryFields = []string{"Id", "Name", "Level", "ParentId"}
)

type AdminCategoryViewData struct {
	CategoryInput *duoerlapi.CategoryInput
	ApiCategories []*duoerlapi.Category
	Validated     *govalidations.Validated
}

// ----------------

func Index(env Env) (status Status, headers Headers, body Body) {

	viewData := &AdminCategoryViewData{
		CategoryInput: new(duoerlapi.CategoryInput),
		ApiCategories: services.GetFullCategories(),
	}

	mangotemplate.ForRender(env, "admin/categories", viewData)
	return
}

func Create(env Env) (status Status, headers Headers, body Body) {
	categoryInput := new(duoerlapi.CategoryInput)
	formdata.UnmarshalByNames(env.Request().Request, &categoryInput, categoryFields)

	viewData := &AdminCategoryViewData{
		CategoryInput: categoryInput,
		ApiCategories: services.GetFullCategories(),
	}

	_, err := services.CreateCategory(categoryInput)
	if validated, ok := err.(*govalidations.Validated); ok {
		viewData.Validated = validated
		mangotemplate.ForRender(env, "admin/categories", viewData)
		return
	}

	if err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/admin/categories")
}
