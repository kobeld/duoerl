package accounts

import (
	"github.com/kobeld/duoerl/models/accounts"
	"github.com/kobeld/duoerl/services"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/formdata"
	"github.com/sunfmin/mangotemplate"
	"labix.org/v2/mgo/bson"
	"net/http"
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
	mangotemplate.ForRender(env, "accounts/edit_profile", &TemplateData{Account: account})
	return
}

func EditProfileAction(env Env) (status Status, headers Headers, body Body) {
	account := services.FetchAccountFromEnv(env)
	formdata.UnmarshalByNames(env.Request().Request, account,
		[]string{"Profile.Gender", "Profile.Location", "Profile.Description", "Profile.HairTexture"})
	if err := account.Save(); err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/profile/edit")
}
