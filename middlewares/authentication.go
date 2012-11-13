package middlewares

import (
	"github.com/kobeld/duoerl/models/accounts"
	sSessions "github.com/kobeld/duoerl/services/sessions"
	. "github.com/paulbellamy/mango"
	"labix.org/v2/mgo/bson"
)

func AuthenticateAccount() Middleware {
	return func(env Env, app App) (status Status, headers Headers, body Body) {

		accountId, exist := sSessions.FetchAccountIdFromSession(env)
		if exist {
			if account, _ := accounts.FindById(bson.ObjectIdHex(accountId)); account != nil {
				sSessions.PutAccountToSession(env, account)
			}
		}

		return app(env)
	}
}
