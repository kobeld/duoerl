package sessions

import (
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerl/services"
	. "github.com/paulbellamy/mango"
	"github.com/theplant/formdata"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"net/http"
)

type SessionData struct {
	User      *users.User
	Validated *govalidations.Validated
}

func LoginPage(env Env) (status Status, headers Headers, body Body) {
	user := services.FetchUserFromEnv(env)
	if user != nil {
		return Redirect(http.StatusFound, "/")
	}
	mangotemplate.ForRender(env, "sessions/login", &SessionData{User: users.NewUser()})
	return
}

func LoginAction(env Env) (status Status, headers Headers, body Body) {
	user := users.NewUser()
	formdata.UnmarshalByNames(env.Request().Request, user, []string{"Email", "Password"})

	validated := user.ValidateLoginForm()
	if validated.HasError() {
		mangotemplate.ForRender(env, "sessions/login", &SessionData{User: user, Validated: validated})
		return
	}

	loginUser := users.LoginWith(user.Email, user.Password)
	if loginUser == nil {
		validated.AddError("Password", "User and password do not match!")
		mangotemplate.ForRender(env, "sessions/login", &SessionData{User: user, Validated: validated})
		return
	}

	services.PutUserIdToSession(env, loginUser.Id.Hex())
	return Redirect(http.StatusFound, "/")
}

func SignupPage(env Env) (status Status, headers Headers, body Body) {
	user := services.FetchUserFromEnv(env)
	if user != nil {
		return Redirect(http.StatusFound, "/")
	}
	mangotemplate.ForRender(env, "sessions/signup", &SessionData{User: users.NewUser()})

	return
}

func SignupAction(env Env) (status Status, headers Headers, body Body) {

	user := users.NewUser()
	formdata.UnmarshalByNames(env.Request().Request, user,
		[]string{"Email", "Name", "Password", "ConfirmPassword"})

	if validated := user.ValidateSignupForm(); validated.HasError() {
		mangotemplate.ForRender(env, "sessions/signup", &SessionData{User: user, Validated: validated})
		return
	}

	if validated := user.ValidateEmailExist(); validated.HasError() {
		mangotemplate.ForRender(env, "sessions/signup", &SessionData{User: user, Validated: validated})
		return
	}

	if err := user.Signup(); err != nil {
		panic(err)
		return
	}

	services.PutUserIdToSession(env, user.Id.Hex())

	return Redirect(http.StatusFound, "/")
}

func Logout(env Env) (status Status, headers Headers, body Body) {
	services.DeleteUserInSession(env)
	return Redirect(http.StatusFound, "/login")
}
