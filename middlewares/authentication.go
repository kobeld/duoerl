package middlewares

import (
	"github.com/kobeld/duoerl/models/accounts"
	"github.com/kobeld/duoerl/services"
	. "github.com/paulbellamy/mango"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func AuthenticateAccount() Middleware {
	return func(env Env, app App) (status Status, headers Headers, body Body) {
		accountId := services.FetchAccountIdFromSession(env)
		if accountId != "" {
			if account, _ := accounts.FindById(bson.ObjectIdHex(accountId)); account != nil {
				services.PutAccountToEnv(env, account)
			}
		}

		return app(env)
	}
}

func HardAuthenAccount() Middleware {
	return func(env Env, app App) (status Status, headers Headers, body Body) {
		accountId := services.FetchAccountIdFromSession(env)
		if accountId == "" {
			return Redirect(http.StatusFound, "/logout")
		}

		account, _ := accounts.FindById(bson.ObjectIdHex(accountId))
		if account == nil {
			return Redirect(http.StatusFound, "/logout")
		}

		services.PutAccountToEnv(env, account)
		return app(env)
	}
}
