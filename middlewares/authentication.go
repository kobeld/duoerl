package middlewares

import (
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerl/services"
	. "github.com/paulbellamy/mango"
	"labix.org/v2/mgo/bson"
	"net/http"
)

func AuthenticateUser() Middleware {
	return func(env Env, app App) (status Status, headers Headers, body Body) {
		userId := services.FetchUserIdFromSession(env)
		if userId != "" {
			if user, _ := users.FindById(bson.ObjectIdHex(userId)); user != nil {
				services.PutUserToEnv(env, user)
			}
		}

		return app(env)
	}
}

func HardAuthenUser() Middleware {
	return func(env Env, app App) (status Status, headers Headers, body Body) {
		userId := services.FetchUserIdFromSession(env)
		if userId == "" {
			return Redirect(http.StatusFound, "/logout")
		}

		user, _ := users.FindById(bson.ObjectIdHex(userId))
		if user == nil {
			return Redirect(http.StatusFound, "/logout")
		}

		services.PutUserToEnv(env, user)
		return app(env)
	}
}
