package reviews

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
	reviewFields = []string{"Id", "ProductId", "Content"}
)

func Create(env Env) (status Status, headers Headers, body Body) {

	reviewInput := new(duoerlapi.ReviewInput)
	formdata.UnmarshalByNames(env.Request().Request, &reviewInput, reviewFields)
	reviewInput.AuthorId = services.FetchAccountIdFromSession(env)

	result, err := services.CreateReview(reviewInput)
	if validated, ok := err.(*govalidations.Validated); ok {
		// TODO: Ajax needed here
		mangotemplate.ForRender(env, "products/new", validated)
		return
	}
	if err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/product/"+result.ProductId)
}
