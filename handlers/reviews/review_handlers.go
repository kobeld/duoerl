package reviews

import (
	"encoding/json"
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"github.com/theplant/formdata"
	"net/http"
)

var (
	reviewFields = []string{"Id", "ProductId", "Content", "Rating", "EfficacyIds"}
)

type ReviewViewData struct {
	LikeCount int
	Validated *govalidations.Validated
}

func Create(env Env) (status Status, headers Headers, body Body) {

	reviewInput := new(duoerlapi.ReviewInput)
	formdata.UnmarshalByNames(env.Request().Request, &reviewInput, reviewFields)
	reviewInput.AuthorId = services.FetchUserIdFromSession(env)

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

func Like(env Env) (status Status, headers Headers, body Body) {
	reviewIdHex := env.Request().FormValue("rid")
	userIdHex := services.FetchUserIdFromSession(env)

	count, err := services.LikeReview(userIdHex, reviewIdHex)
	viewData := &ReviewViewData{LikeCount: count}
	if validated, ok := err.(*govalidations.Validated); ok {
		viewData.Validated = validated
		b, _ := json.Marshal(viewData)
		body = Body(b)
		return
	}

	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(viewData)
	body = Body(b)

	return
}
