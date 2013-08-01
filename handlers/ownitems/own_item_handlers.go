package ownitems

import (
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/formdata"
)

var (
	ownItemFields = []string{"ProductId", "GotFrom"}
)

func Create(env Env) (status Status, headers Headers, body Body) {

	ownItemInput := new(duoerlapi.OwnItemInput)
	formdata.UnmarshalByNames(env.Request().Request, &ownItemInput, ownItemFields)

	ownItemInput.UserId = services.FetchUserIdFromSession(env)

	err := services.AddOwnItem(ownItemInput)
	if err != nil {
		panic(err)
	}

	return
}

func Delete(env Env) (status Status, headers Headers, body Body) {

	// productId := env.Request().FormValue("pid")
	// userId := services.FetchUserIdFromSession(env)

	// err := services.RemoveOwnItem(userId, productId)
	// if err != nil {
	// 	panic(err)
	// }

	return
}
