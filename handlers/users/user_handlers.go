package users

import (
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/mangotemplate"
	"github.com/theplant/formdata"
	"labix.org/v2/mgo/bson"
	"net/http"
)

var (
	userFields = []string{"Avatar", "Profile.Gender", "Profile.Location", "Profile.Description",
		"Profile.HairTexture", "Profile.SkinTexture", "Profile.Birthday"}
)

type UserViewData struct {
	ApiUser            *duoerlapi.User
	ApiNotes           []*duoerlapi.Note
	ApiPosts           []*duoerlapi.Post
	NewPostId          string
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

	apiNotes, err := services.GetUserNotes(id)
	if err != nil {
		panic(err)
	}

	apiPosts, err := services.GetUserPosts(id)
	if err != nil {
		panic(err)
	}

	userViewData := &UserViewData{
		ApiUser:   apiUser,
		ApiNotes:  apiNotes,
		ApiPosts:  apiPosts,
		NewPostId: bson.NewObjectId().Hex(),
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

	userInput := &duoerlapi.UserInput{Id: services.FetchUserIdFromSession(env)}
	formdata.UnmarshalByNames(env.Request().Request, userInput, userFields)

	err := services.UpdateProfile(userInput)
	if err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/user/edit")
}
