package users

import (
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/formdata"
	"github.com/sunfmin/mangotemplate"
	"net/http"
)

var (
	userFields = []string{"Profile.Gender", "Profile.Location", "Profile.Description",
		"Profile.HairTexture", "Profile.SkinTexture"}
)

type UserViewData struct {
	ApiUser            *duoerlapi.User
	IsCurrent          bool
	SkinTextureOptions map[string]string
	HairTextureOptions map[string]string
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

	userViewData := &UserViewData{
		ApiUser:            apiUser,
		SkinTextureOptions: global.SkinTextureOptions,
		HairTextureOptions: global.HairTextureOptions,
	}

	mangotemplate.ForRender(env, "users/edit", userViewData)
	return
}

func Update(env Env) (status Status, headers Headers, body Body) {
	user := services.FetchUserFromEnv(env)
	formdata.UnmarshalByNames(env.Request().Request, user, userFields)

	if err := user.Save(); err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/user/edit")
}
