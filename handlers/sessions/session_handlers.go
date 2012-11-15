package sessions

import (
	"github.com/kobeld/duoerl/models/accounts"
	"github.com/kobeld/duoerl/services"
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
	account := services.FetchAccountFromEnv(env)
	if account != nil {
		return Redirect(http.StatusFound, "/")
	}
	mangotemplate.ForRender(env, "sessions/login", &SessionData{Account: accounts.NewAccount()})
	return
}

func LoginAction(env Env) (status Status, headers Headers, body Body) {
	account := accounts.NewAccount()
	formdata.UnmarshalByNames(env.Request().Request, account, []string{"Email", "Password"})

	validated := account.ValidateLoginForm()
	if validated.HasError() {
		mangotemplate.ForRender(env, "sessions/login", &SessionData{Account: account, Validated: validated})
		return
	}

	loginAccount := accounts.LoginWith(account.Email, account.Password)
	if loginAccount == nil {
		validated.AddError("Password", "Account and password do not match!")
		mangotemplate.ForRender(env, "sessions/login", &SessionData{Account: account, Validated: validated})
		return
	}

	services.PutAccountIdToSession(env, loginAccount.Id.Hex())
	return Redirect(http.StatusFound, "/")
}

func SignupPage(env Env) (status Status, headers Headers, body Body) {
	account := services.FetchAccountFromEnv(env)
	if account != nil {
		return Redirect(http.StatusFound, "/")
	}
	mangotemplate.ForRender(env, "sessions/signup", &SessionData{Account: accounts.NewAccount()})

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

	services.PutAccountIdToSession(env, account.Id.Hex())

	return Redirect(http.StatusFound, "/")
}

func Logout(env Env) (status Status, headers Headers, body Body) {
	services.DeleteAccountInSession(env)
	return Redirect(http.StatusFound, "/login")
}
