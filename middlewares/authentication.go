package middlewares

import (
	"github.com/kobeld/duoerl/models/accounts"
	sSessions "github.com/kobeld/duoerl/services/sessions"
	. "github.com/paulbellamy/mango"
	"labix.org/v2/mgo/bson"
)

func AuthenticateAccount() Middleware {
	return func(env Env, app App) (status Status, headers Headers, body Body) {
		accountId := sSessions.FetchAccountIdFromSession(env)
		if accountId != "" {
			if account, _ := accounts.FindById(bson.ObjectIdHex(accountId)); account != nil {
				sSessions.PutAccountToEnv(env, account)
			}
		}

		return app(env)
	}
}
