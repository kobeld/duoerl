package posts

import (
	"encoding/json"
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/govalidations"
	"github.com/theplant/formdata"
)

var (
	postFields = []string{"Id", "Content"}
)

func Create(env Env) (status Status, headers Headers, body Body) {

	postInput := new(duoerlapi.PostInput)
	formdata.UnmarshalByNames(env.Request().Request, &postInput, postFields)
	postInput.AuthorId = services.FetchUserIdFromSession(env)

	result, err := services.CreatePost(postInput)
	if validated, ok := err.(*govalidations.Validated); ok {
		b, _ := json.Marshal(validated)
		body = Body(b)
		return
	}
	if err != nil {
		panic(err)
	}

	b, _ := json.Marshal(result)
	body = Body(b)
	return
}
