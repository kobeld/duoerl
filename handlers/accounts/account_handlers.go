package accounts

import (
	"github.com/kobeld/duoerl/models/accounts"
	"github.com/kobeld/duoerl/services"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/mangotemplate"
	"labix.org/v2/mgo/bson"
)

type TemplateData struct {
	Account          *accounts.Account
	IsCurrentAccount bool
}

func ShowProfile(env Env) (status Status, headers Headers, body Body) {
	id := env.Request().URL.Query().Get(":id")

	account, err := accounts.FindById(bson.ObjectIdHex(id))
	if err != nil {
		panic(err)
	}

	isCurrent := services.IsCurrentAccountWithId(env, id)

	mangotemplate.ForRender(env, "accounts/profile",
		&TemplateData{Account: account, IsCurrentAccount: isCurrent})
	return
}

func EditProfile(env Env) (status Status, headers Headers, body Body) {
	account := services.FetchAccountFromEnv(env)
	if account == nil {
		status = 500
		return
	}

	mangotemplate.ForRender(env, "accounts/edit_profile", &TemplateData{Account: account})
	return
}

func EditProfileAction(env Env) (status Status, headers Headers, body Body) {
	return
}
