package users

import (
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/formdata"
	"github.com/sunfmin/mangotemplate"
	"net/http"
)

type UserViewData struct {
	ApiUser   *duoerlapi.User
	IsCurrent bool
}

// ------------------

func Show(env Env) (status Status, headers Headers, body Body) {
	id := env.Request().URL.Query().Get(":id")

	apiUser, err := services.GetUser(id)
	if err != nil {
		panic(err)
	}

	userViewData := &UserViewData{
		ApiUser:   apiUser,
		IsCurrent: services.IsCurrentUserWithId(env, id),
	}

	mangotemplate.ForRender(env, "users/show", userViewData)
	return
}

func Edit(env Env) (status Status, headers Headers, body Body) {
	id := services.FetchUserIdFromSession(env)

	apiUser, err := services.GetUser(id)
	if err != nil {
		panic(err)
	}

	mangotemplate.ForRender(env, "users/edit", &UserViewData{ApiUser: apiUser})
	return
}

func Update(env Env) (status Status, headers Headers, body Body) {
	account := services.FetchUserFromEnv(env)
	formdata.UnmarshalByNames(env.Request().Request, account,
		[]string{"Profile.Gender", "Profile.Location", "Profile.Description", "Profile.HairTexture"})
	if err := account.Save(); err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/profile/edit")
}
