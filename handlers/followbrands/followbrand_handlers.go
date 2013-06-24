package followbrands

import (
	"github.com/kobeld/duoerl/services"
	. "github.com/paulbellamy/mango"
)

func Create(env Env) (status Status, headers Headers, body Body) {

	brandId := env.Request().FormValue("bid")
	userId := services.FetchAccountIdFromSession(env)

	err := services.CreateFollowBrand(userId, brandId)
	if err != nil {
		panic(err)
	}

	return
}

func Delete(env Env) (status Status, headers Headers, body Body) {

	brandId := env.Request().FormValue("bid")
	userId := services.FetchAccountIdFromSession(env)

	err := services.DeleteFollowBrand(userId, brandId)
	if err != nil {
		panic(err)
	}

	return
}
