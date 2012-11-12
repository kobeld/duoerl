package sessions

import (
	"github.com/kobeld/duoerl/models/accounts"
	sSessions "github.com/kobeld/duoerl/services/sessions"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/formdata"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"net/http"
)

type SessionData struct {
	Account   *accounts.Account
	Validated *govalidations.Validated
}

func LoginPage(env Env) (status Status, headers Headers, body Body) {
	mangotemplate.ForRender(env, "sessions/login", &SessionData{Account: accounts.NewAccount()})
	return
}

func LoginAction(env Env) (status Status, headers Headers, body Body) {
	account := accounts.NewAccount()
	formdata.UnmarshalByNames(env.Request().Request, account, []string{"Email", "Password"})

	if validated := account.ValidateLoginForm(); validated.HasError() {
		mangotemplate.ForRender(env, "sessions/login", &SessionData{Account: account, Validated: validated})
		return
	}

	if validated := account.ValidateLoginAccount(); validated.HasError() {
		mangotemplate.ForRender(env, "sessions/login", &SessionData{Account: account, Validated: validated})
		return
	}

	sSessions.PutUserIdToSession(env, account.Id.Hex())

	return Redirect(http.StatusFound, "/")
}

func SignupPage(env Env) (status Status, headers Headers, body Body) {
	account := new(accounts.Account)
	mangotemplate.ForRender(env, "sessions/signup", &SessionData{Account: account})

	return
}

func SignupAction(env Env) (status Status, headers Headers, body Body) {

	account := accounts.NewAccount()
	formdata.UnmarshalByNames(env.Request().Request, account,
		[]string{"Email", "Name", "Password", "ConfirmPassword"})

	if validated := account.ValidateSignupForm(); validated.HasError() {
		mangotemplate.ForRender(env, "sessions/signup", &SessionData{Account: account, Validated: validated})
		return
	}

	if validated := account.ValidateEmailExist(); validated.HasError() {
		mangotemplate.ForRender(env, "sessions/signup", &SessionData{Account: account, Validated: validated})
		return
	}

	if err := account.Signup(); err != nil {
		panic(err)
		return
	}

	sSessions.PutUserIdToSession(env, account.Id.Hex())

	return Redirect(http.StatusFound, "/")
}
