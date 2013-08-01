package wishitems

import (
	"github.com/kobeld/duoerl/services"
	. "github.com/paulbellamy/mango"
)

func Create(env Env) (status Status, headers Headers, body Body) {

	productId := env.Request().FormValue("pid")
	userId := services.FetchUserIdFromSession(env)

	err := services.AddWishItem(userId, productId)
	if err != nil {
		panic(err)
	}

	return
}

func Delete(env Env) (status Status, headers Headers, body Body) {

	productId := env.Request().FormValue("pid")
	userId := services.FetchUserIdFromSession(env)

	err := services.RemoveWishItem(userId, productId)
	if err != nil {
		panic(err)
	}

	return
}
