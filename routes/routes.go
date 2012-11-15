package routes

import (
	"github.com/bmizerany/pat"
	"github.com/kobeld/duoerl/handlers/accounts"
	"github.com/kobeld/duoerl/handlers/feeds"
	"github.com/kobeld/duoerl/handlers/sessions"
	"github.com/kobeld/duoerl/middlewares"
	"github.com/kobeld/mangogzip"
	"github.com/paulbellamy/mango"
	"github.com/sunfmin/mangolog"
	"net/http"
)

func Mux() (mux *http.ServeMux) {
	p := pat.New()
	sessionMW := mango.Sessions("f908b1c425062e95d30b8d30de7123458", "duoerl",
		&mango.CookieOptions{Path: "/", MaxAge: 3600 * 24 * 7})
	rendererMW := middlewares.ProduceRenderer()
	authenMW := middlewares.AuthenticateAccount()
	rHtml, _ := middlewares.RespondHtml(), middlewares.RespondJson()

	mainLayoutMW := middlewares.ProduceLayout(middlewares.MAIN_LAYOUT)
	mainStack := new(mango.Stack)
	mainStack.Middleware(mangogzip.Zipper, mangolog.Logger, sessionMW, authenMW, mainLayoutMW, rendererMW, rHtml)

	p.Get("/login", mainStack.HandlerFunc(sessions.LoginPage))
	p.Post("/login", mainStack.HandlerFunc(sessions.LoginAction))
	p.Get("/signup", mainStack.HandlerFunc(sessions.SignupPage))
	p.Post("/signup", mainStack.HandlerFunc(sessions.SignupAction))
	p.Get("/logout", mainStack.HandlerFunc(sessions.Logout))

	p.Get("/profile/edit", mainStack.HandlerFunc(accounts.EditProfile))
	p.Get("/profile/:id", mainStack.HandlerFunc(accounts.ShowProfile))

	p.Get("/", mainStack.HandlerFunc(feeds.Index))

	mux = http.NewServeMux()
	mux.HandleFunc("/favicon.ico", filterUrl)
	mux.Handle("/", p)
	mux.Handle("/public/", http.FileServer(http.Dir(".")))
	return
}

func filterUrl(w http.ResponseWriter, r *http.Request) {
	return
}
