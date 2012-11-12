package sessions

import (
	"github.com/kobeld/duoerl/models/accounts"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/formdata"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"net/http"
)

type SessionData struct {
	Account *accounts.Account
	Errors  govalidations.Errors
}

func LoginPage(env Env) (status Status, headers Headers, body Body) {
	mangotemplate.ForRender(env, "sessions/login", &SessionData{Account: accounts.NewAccount()})
	return
}

func LoginAction(env Env) (status Status, headers Headers, body Body) {
	account := accounts.NewAccount()
	formdata.UnmarshalByNames(env.Request().Request, account, []string{"Email", "Name"})

	errors := account.ValidateLogin()
	if errors != nil {
		mangotemplate.ForRender(env, "sessions/signup", &SessionData{Account: account, Errors: errors})
		return
	}

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

	errors := account.ValidateSignup()
	if errors != nil {
		mangotemplate.ForRender(env, "sessions/signup", &SessionData{Account: account, Errors: errors})
		return
	}

	return Redirect(http.StatusFound, "/")
}
