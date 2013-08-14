package efficacies

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
	efficacyFields = []string{"Id", "Name", "ParentId"}
)

type AdminEfficacyViewData struct {
	EfficacyInput *duoerlapi.EfficacyInput
	ApiCategories []*duoerlapi.Category
	Validated     *govalidations.Validated
}

// ----------------

func Create(env Env) (status Status, headers Headers, body Body) {

	efficacyInput := new(duoerlapi.EfficacyInput)
	formdata.UnmarshalByNames(env.Request().Request, &efficacyInput, efficacyFields)
	viewData := &AdminEfficacyViewData{
		EfficacyInput: efficacyInput,
		ApiCategories: services.GetFullCategories(),
	}

	_, err := services.CreateEfficacy(efficacyInput)
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
